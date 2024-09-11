package todo

import (
	"errors"
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
