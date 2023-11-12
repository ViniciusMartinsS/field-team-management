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

const (
	badRequestMessage     = "malformed request"
	internalServerMessage = "internal server error"
	forbiddenMessage      = "not allowed to perform this action"
)

type taskCreateRequest struct {
	Summary string `json:"summary" binding:"required,max=2500"`
	Date    string `json:"date"`
	UserID  int    `json:"user_id" binding:"required"`
}

type taskUpdateRequest struct {
	Summary string `json:"summary" binding:"max=2500"`
	Date    string `json:"date"`
}

type taskResponse struct {
	ID      int    `json:"id"`
	Summary string `json:"summary"`
	Date    string `json:"date"`
	UserID  int    `json:"user_id"`
}

type taskHandlerResponse struct {
	Status bool   `json:"status"`
	Result any    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

type TaskAPIHandler struct {
	router        *gin.Engine
	authenticator domain.Authenticator
	taskUsecase   domain.TaskUsecase
}

func NewTask(r *gin.Engine, authenticator domain.Authenticator, taskUsecase domain.TaskUsecase) (*TaskAPIHandler, error) {
	if r == nil {
		return &TaskAPIHandler{}, errors.New("router must not be nil")
	}

	if authenticator == nil {
		return &TaskAPIHandler{}, errors.New("authenticator must not be nil")
	}

	if taskUsecase == nil {
		return &TaskAPIHandler{}, errors.New("taskUsecase must not be nil")
	}

	return &TaskAPIHandler{
		router:        r,
		authenticator: authenticator,
		taskUsecase:   taskUsecase,
	}, nil
}

func (h *TaskAPIHandler) CreateRouter() {
	v1 := h.router.Group("v1")
	tasks := v1.Group("tasks")
	tasks.Use(Authenticator(h.authenticator))

	tasks.GET("", h.get)
	tasks.POST("", h.post)
	tasks.PATCH("/:id", h.patch)
	tasks.DELETE("/:id", h.remove)
}

func (h *TaskAPIHandler) get(c *gin.Context) {
	user := identifyUserRequester(c)

	result, err := h.taskUsecase.ListByUser(context.Background(), user)
	if err != nil {
		if errors.Is(err, domain.ErrTasksNotFound) {
			c.JSON(http.StatusBadRequest, toResponse(false, err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, toResponse(false, internalServerMessage))
		return
	}

	c.JSON(http.StatusOK, toResponse(true, formatResponseSlice(result)))
}

func (h *TaskAPIHandler) post(c *gin.Context) {
	var request taskCreateRequest

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, toResponse(false, badRequestMessage))
		return
	}

	parsedDate, err := parseDate(request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, toResponse(false, badRequestMessage))
		return
	}

	task, err := domain.NewTask(request.Summary, parsedDate, request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, toResponse(false, badRequestMessage))
		return
	}

	user := identifyUserRequester(c)

	result, err := h.taskUsecase.Add(context.Background(), task, user)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotAllowed) {
			c.JSON(http.StatusForbidden, toResponse(false, forbiddenMessage))
			return
		}

		c.JSON(http.StatusInternalServerError, toResponse(false, internalServerMessage))
		return
	}

	c.JSON(http.StatusOK, toResponse(true, formatResponseSingle(result)))
}

func (h *TaskAPIHandler) patch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, toResponse(false, badRequestMessage))
		return
	}

	var request taskUpdateRequest

	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, toResponse(false, badRequestMessage))
		return
	}

	parsedDate, err := parseDate(request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, toResponse(false, badRequestMessage))
		return
	}

	task := domain.Task{
		ID:      id,
		Summary: request.Summary,
		Date:    parsedDate,
	}

	user := identifyUserRequester(c)

	result, err := h.taskUsecase.Update(context.Background(), task, user)
	if err != nil {
		if errors.Is(err, domain.ErrTasksNotFound) {
			c.JSON(http.StatusBadRequest, toResponse(false, err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, toResponse(false, internalServerMessage))
		return
	}

	c.JSON(http.StatusOK, toResponse(true, formatResponseSingle(result)))
}

func (h *TaskAPIHandler) remove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, toResponse(false, badRequestMessage))
		return
	}

	user := identifyUserRequester(c)

	err = h.taskUsecase.Remove(context.Background(), id, user)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotAllowed) {
			c.JSON(http.StatusForbidden, toResponse(false, forbiddenMessage))
			return
		}

		c.JSON(http.StatusInternalServerError, toResponse(false, internalServerMessage))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func identifyUserRequester(c *gin.Context) domain.User {
	userID := c.MustGet("user_id").(float64)
	roleID := c.MustGet("role_id").(float64)

	return domain.User{
		ID:     int(userID),
		RoleID: int(roleID),
	}
}

func toResponse(status bool, result any) taskHandlerResponse {
	switch status {
	case true:
		return taskHandlerResponse{
			Status: status,
			Result: result,
		}
	case false:
		return taskHandlerResponse{
			Status: false,
			Error:  result.(string),
		}
	}

	return taskHandlerResponse{}
}

func formatResponseSingle(task domain.Task) taskResponse {
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

func formatResponseSlice(task []domain.Task) []taskResponse {
	var response []taskResponse

	for _, t := range task {
		response = append(response, formatResponseSingle(t))
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
