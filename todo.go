package main

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

func (todos *Todos) addTask(title string) {
	todo := Task{
		Title:       title,
		Status:      TODO,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(idx int) error {
	if idx < 0 || idx >= len(*todos) {
		err := errors.New("Invalid index")
		return err
	}
	return nil
}

func (todos *Todos) removeTask(idx int) error {
	t := *todos
	if err := t.validateIndex(idx); err != nil {
		return err
	}

	*todos = append(t[:idx], t[idx+1:]...)

	return nil
}
