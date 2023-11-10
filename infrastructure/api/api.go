package api

import (
	"context"
	"github.com/ViniciusMartinss/field-team-management/application/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTaskRoutes(r *gin.Engine) {
	tasks := r.Group("tasks")

	tasks.GET("", getTasks())
}

func getTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskUsecase, _ := usecase.NewTask(nil, nil)
		tasks, _ := taskUsecase.ListByUserID(context.Background(), 1)

		c.JSON(http.StatusOK, tasks)
	}
}
