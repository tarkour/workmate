package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/tarkour/workmate/internal/models"
)

func CreateTask() *models.Task {

	store := GetTaskStore()
	store.Lock()
	defer store.Unlock()

	taskID := uuid.New().String()
	newTask := &models.Task{
		ID:        taskID,
		Status:    "just created",
		CreatedAt: time.Now(),
	}
	store.Tasks[taskID] = newTask

	return newTask
}

func GetTask(taskID string) (*models.Task, bool) {

	store := GetTaskStore()
	store.RLock()
	defer store.RUnlock()
	task, exists := store.Tasks[taskID]

	return task, exists
}

func DeleteTask(taskID string) bool {
	store := GetTaskStore()
	store.Lock()
	defer store.Unlock()

	if _, exists := store.Tasks[taskID]; !exists {
		return false
	}

	delete(store.Tasks, taskID)
	return true
}
