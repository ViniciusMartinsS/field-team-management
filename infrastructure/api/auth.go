package api

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const unauthorizedMessage = "unauthorized"

type authRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthAPIHandler struct {
	router      *gin.Engine
	authUsecase domain.AuthUsecase
}

func NewAuth(r *gin.Engine, authUsecase domain.AuthUsecase) (*AuthAPIHandler, error) {
	if r == nil {
		return &AuthAPIHandler{}, errors.New("router must not be nil")
	}

	if authUsecase == nil {
		return &AuthAPIHandler{}, errors.New("authUsecase must not be nil")
	}

	return &AuthAPIHandler{
		router:      r,
		authUsecase: authUsecase,
	}, nil
}

func (h *AuthAPIHandler) CreateRouter() {
	v1 := h.router.Group("v1")
	tasks := v1.Group("auth")

	tasks.POST("", h.post)
}

func (h *AuthAPIHandler) post(c *gin.Context) {
	var request authRequest

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, toResponse(false, badRequestMessage))
		return
	}

	result, err := h.authUsecase.Authenticate(context.Background(), request.Email, request.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) || errors.Is(err, domain.ErrUserInvalidPass) {
			c.JSON(http.StatusBadRequest, toResponse(false, unauthorizedMessage))
			return
		}

		c.JSON(http.StatusInternalServerError, toResponse(false, internalServerMessage))
		return
	}

	c.JSON(http.StatusOK, toResponse(true, result))
}
