package main

import (
	"github.com/gin-gonic/gin"
	"github.com/chiarorosa/go-todo-api/internal/tasks"
)

func main() {
	r := gin.Default()

	repo := tasks.NewRepository()
	handler := tasks.NewHandler(repo)
	handler.RegisterRoutes(r)

	r.Run(":8080")
}