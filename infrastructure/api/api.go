package api

import (
	"context"
	"github.com/ViniciusMartinss/field-team-management/application/usecase"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/encryption"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTaskRoutes(r *gin.Engine) {
	tasks := r.Group("tasks")

	tasks.GET("", getTasks())
}

func getTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskRepository, _ := repository.NewTask(nil)
		encryptor, _ := encryption.New("123456789123456789123456")
		taskUsecase, _ := usecase.NewTask(taskRepository, taskRepository, nil, encryptor)
		tasks, _ := taskUsecase.ListByUserID(context.Background(), 1)

		c.JSON(http.StatusOK, tasks)
	}
}
