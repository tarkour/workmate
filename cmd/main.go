package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/tasks", postfunc)           // create task
	e.GET("/tasks/{:id}", getfunc)       //get some task
	e.DELETE("/tasks/{:id}", deletefunc) // delete some task

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server start error: %v", err)
	}

	fmt.Println("Server launched at localhost:8080")
}
