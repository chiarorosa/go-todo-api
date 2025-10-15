package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/tasks", h.getTasks)
	r.POST("/tasks", h.createTask)
}

func (h *Handler) getTasks(c *gin.Context) {
	tasks := h.repo.GetAll()
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	created := h.repo.Create(task)
	c.JSON(http.StatusCreated, created)
}