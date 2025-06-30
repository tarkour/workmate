package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tarkour/workmate/api/process"
	"github.com/tarkour/workmate/internal/storage"
)

func CreateTaskHandler(c echo.Context) error {
	newTask := storage.CreateTask()
	go process.ProcessImitation(newTask.ID)

	return c.JSON(http.StatusOK, map[string]string{
		"id": newTask.ID,
	})
}

func GetTaskHandler(c echo.Context) error {
	taskID := c.Param("id")
	task, exists := storage.GetTask(taskID)
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "task not found",
		})
	}
	return c.JSON(http.StatusOK, task)
}

func DeleteTaskHandler(c echo.Context) error {
	taskID := c.Param("id")
	if exists := storage.DeleteTask(taskID); !exists {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "task not found",
		})
	}

	return c.NoContent(http.StatusNoContent)

}
