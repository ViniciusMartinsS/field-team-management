package api

import (
	"context"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskAPIHandler struct {
	router      *gin.Engine
	taskUsecase domain.TaskUsecase
}

func NewTask(r *gin.Engine, taskUsecase domain.TaskUsecase) *TaskAPIHandler {
	return &TaskAPIHandler{
		router:      r,
		taskUsecase: taskUsecase,
	}
}

func (h *TaskAPIHandler) CreateRouter() {
	v1 := h.router.Group("v1")
	tasks := v1.Group("tasks")

	tasks.GET("", h.get)
	tasks.POST("", h.post)
	tasks.PATCH("", h.patch)
	tasks.DELETE("/:id", h.remove)
}

func (h *TaskAPIHandler) get(c *gin.Context) {
	tasks, _ := h.taskUsecase.ListByUserID(context.Background(), 3)

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskAPIHandler) post(c *gin.Context) {
	tasks, _ := h.taskUsecase.Add(context.Background(), domain.Task{
		Summary: "This is User Task 02",
		UserID:  2,
	})

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskAPIHandler) patch(c *gin.Context) {
	tasks, _ := h.taskUsecase.Update(context.Background(), domain.Task{
		ID:      2,
		Summary: "Be different",
		Date:    nil,
		UserID:  3,
	})

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskAPIHandler) remove(c *gin.Context) {
	err := h.taskUsecase.Remove(context.Background(), 1, 1)

	c.JSON(http.StatusNoContent, err)
}
