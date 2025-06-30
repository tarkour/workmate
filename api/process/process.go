package process

import (
	"math/rand"
	"time"

	"github.com/tarkour/workmate/internal/storage"
)

func ProcessImitation(taskID string) {

	store := storage.GetTaskStore()

	store.Lock()
	task, exists := store.Tasks[taskID]
	if !exists {
		store.Unlock()
		return
	}
	task.Status = "processing"
	task.StartedAt = time.Now()
	store.Unlock()

	minDur := 3 * time.Minute
	maxDur := 3 * time.Minute
	duration := minDur + time.Duration(rand.Int63n(int64(maxDur-minDur)))

	time.Sleep(duration)

	store.Lock()
	defer store.Unlock()

	if task, exists := store.Tasks[taskID]; exists {
		task.Status = "completed"
		task.FinishedAt = time.Now()
		task.Duration = task.FinishedAt.Sub(task.StartedAt)
		task.Result = "Operation done successfuly"
	}
}
