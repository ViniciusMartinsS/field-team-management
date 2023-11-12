package api

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"
)

type taskResponse struct {
	ID      int    `json:"id"`
	Summary string `json:"summary"`
	Date    string `json:"date"`
	UserID  int    `json:"user_id"`
}

type taskCreateRequest struct {
	Summary string `json:"summary" binding:"required"`
	Date    string `json:"date"`
	UserID  int    `json:"user_id" binding:"required"`
}

type taskUpdateRequest struct {
	ID      int    `json:"id" binding:"required"`
	Summary string `json:"summary"`
	Date    string `json:"date"`
	UserID  int    `json:"user_id" binding:"required"`
}

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
	// GET USER ID FROM TOKEN
	result, err := h.taskUsecase.ListByUserID(context.Background(), 2)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusBadRequest, "user not found")
			return
		}

		if errors.Is(err, domain.ErrTasksNotFound) {
			c.JSON(http.StatusBadRequest, "task not found")
			return
		}

		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, toResponseSlice(result))
}

func (h *TaskAPIHandler) post(c *gin.Context) {
	var request taskCreateRequest

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, "required fields were not sent or with invalid content")
		return
	}

	parsedDate, err := parseDate(request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid date format")
		return
	}

	task, err := domain.NewTask(request.Summary, parsedDate, request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.taskUsecase.Add(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, toResponseSingle(result))
}

func (h *TaskAPIHandler) patch(c *gin.Context) {
	var request taskUpdateRequest

	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, "required fields were not sent or with invalid content")
		return
	}

	parsedDate, err := parseDate(request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid date format")
		return
	}

	task := domain.Task{
		ID:      request.ID,
		Summary: request.Summary,
		Date:    parsedDate,
		UserID:  request.UserID,
	}

	result, err := h.taskUsecase.Update(context.Background(), task)
	if err != nil {
		if errors.Is(err, domain.ErrTasksNotFound) {
			c.JSON(http.StatusBadRequest, "task not found")
			return
		}

		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, toResponseSingle(result))
}

func (h *TaskAPIHandler) remove(c *gin.Context) {
	// GET USER ID FROM TOKEN

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err = h.taskUsecase.Remove(context.Background(), id, 1)

	c.JSON(http.StatusNoContent, err)
}

func toResponseSingle(task domain.Task) taskResponse {
	var date string

	if task.Date != nil {
		date = task.Date.Format("2006/01/02 15:04")
	}

	return taskResponse{
		ID:      task.ID,
		Summary: task.Summary,
		Date:    date,
		UserID:  task.UserID,
	}
}

func toResponseSlice(task []domain.Task) []taskResponse {
	var response []taskResponse

	for _, t := range task {
		response = append(response, toResponseSingle(t))
	}

	return response
}

func parseDate(date string) (*time.Time, error) {
	if date == "" {
		return nil, nil
	}

	parsedTime, err := time.Parse("2006/01/02 15:04", date)
	if err != nil {
		return nil, errors.New("bad Request")
	}

	return &parsedTime, nil
}
