package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/tarkour/workmate/api/handlers"
)

func main() {

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.POST("/tasks", handlers.CreateTaskHandler)       // create task
	e.GET("/tasks/:id", handlers.GetTaskHandler)       // get some task
	e.DELETE("/tasks/:id", handlers.DeleteTaskHandler) // delete some task

	fmt.Println("Server launched at localhost:8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server start error: %v", err)
	}

}
