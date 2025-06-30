package models

import (
	"sync"
	"time"
)

type Task struct {
	ID         string
	Status     string
	CreatedAt  time.Time
	StartedAt  time.Time
	FinishedAt time.Time
	Result     string // need?
	Error      error  // need?
}

type TaskStore struct {
	sync.RWMutex
	Task map[string]*Task
}
