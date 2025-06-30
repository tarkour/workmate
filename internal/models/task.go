package models

import (
	"sync"
	"time"
)

type Task struct {
	ID         string        `json:"id"`
	Status     string        `json:"status"`
	CreatedAt  time.Time     `json:"created_at"`
	StartedAt  time.Time     `json:"started_at,omitempty"`
	FinishedAt time.Time     `json:"finished_at,omitempty"`
	Duration   time.Duration `json:"duration,omitempty"`
	Result     string        `json:"result,omitempty"`
	Error      string        `json:"error,omitempty"`
}

type TaskStore struct {
	sync.RWMutex
	Tasks map[string]*Task
}

// internal/models/task.go
func NewTaskStore() *TaskStore {
	return &TaskStore{
		Tasks: make(map[string]*Task, 0),
	}
}
