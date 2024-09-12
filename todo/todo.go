package todo

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

type StatusType int

const (
	TODO StatusType = iota
	DOING
	DONE
)

type Task struct {
	Title       string
	Status      StatusType
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Task

func (todos *Todos) AddTask(title string) (bool, error) {
	todo := Task{
		Title:       title,
		Status:      TODO,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
	return true, nil
}

func (todos *Todos) validateIndex(idx int) error {
	if idx < 0 || idx >= len(*todos) {
		err := errors.New("Invalid index")
		return err
	}
	return nil
}

func (todos *Todos) RemoveTask(idx int) error {
	t := *todos
	if err := t.validateIndex(idx); err != nil {
		return err
	}

	*todos = append(t[:idx], t[idx+1:]...)

	return nil
}

func (todos *Todos) Toggle(idx int) error {
	t := *todos
	if err := t.validateIndex(idx); err != nil {
		return err
	}

	completed := t[idx].Status
	if completed != DONE {
		complTime := time.Now()
		t[idx].CompletedAt = &complTime
		t.updateStatus(idx, DONE)
	}

	return nil
}

func (todos *Todos) updateStatus(idx int, status StatusType) error {
	t := *todos
	if err := t.validateIndex(idx); err != nil {
		return err
	}

	t[idx].Status = status

	return nil
}

func (todos *Todos) UpdateTitle(idx int, title string) error {
	t := *todos
	if err := t.validateIndex(idx); err != nil {
		return err
	}

	t[idx].Title = title

	return nil
}

func (todos *Todos) SaveToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range *todos {
		completedAt := ""
		if task.CompletedAt != nil {
			completedAt = task.CompletedAt.Format("02-Jan-2006 15:04")
		}

		err := writer.Write([]string{
			task.Title,
			strconv.Itoa(int(task.Status)),
			task.CreatedAt.Format("02-Jan-2006 15:04"),
			completedAt,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (todos *Todos) LoadFromCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		if len(records) != 4 {
			return errors.New("invalid record length")
		}

		statusInt, err := strconv.Atoi(record[1])
		if err != nil {
			return err
		}
		statusType := StatusType(statusInt)

		createdAt, err := time.Parse("02-Jan-2006 15:04", record[2])
		if err != nil {
			return err
		}

		var completedAt *time.Time
		if record[3] != "" {
			t, err := time.Parse("02-Jan-2006 15:04", record[3])
			if err != nil {
				return err
			}
			completedAt = &t
		}

		task := Task{
			Title:       record[0],
			Status:      statusType,
			CreatedAt:   createdAt,
			CompletedAt: completedAt,
		}
		*todos = append(*todos, task)
	}

	return nil
}
