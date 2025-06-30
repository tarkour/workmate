package storage

import (
	"sync"

	"github.com/tarkour/workmate/internal/models"
)

var (
	storeInstance *models.TaskStore
	storeOnce     sync.Once
)

func GetTaskStore() *models.TaskStore{
	storeOnce.Do(func() {
		storeInstance = models.NewTaskStore()
	})
	return storeInstance
}