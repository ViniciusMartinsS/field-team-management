package api

import (
	"context"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type authRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthAPIHandler struct {
	router      *gin.Engine
	authUsecase domain.AuthUsecase
}

func NewAuth(r *gin.Engine, authUsecase domain.AuthUsecase) *AuthAPIHandler {
	return &AuthAPIHandler{
		router:      r,
		authUsecase: authUsecase,
	}
}

func (h *AuthAPIHandler) CreateRouter() {
	v1 := h.router.Group("v1")
	tasks := v1.Group("auth")

	tasks.POST("", h.post)
}

func (h *AuthAPIHandler) post(c *gin.Context) {
	var request authRequest

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, "required fields were not sent or with invalid content")
		return
	}

	result, err := h.authUsecase.Authenticate(context.Background(), request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, result)
}
