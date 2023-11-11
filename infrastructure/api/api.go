package api

import (
	"context"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/ViniciusMartinss/field-team-management/application/usecase"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/encryption"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func CreateTaskRoutes(r *gin.Engine, db *sqlx.DB) {
	tasks := r.Group("tasks")

	tasks.GET("", getTasks(db))
	tasks.POST("", postTasks(db))
}

func getTasks(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskRepository, _ := repository.NewTask(db)
		userRepository, _ := repository.NewUser(db)

		encryptor, _ := encryption.New("123456789123456789123456")

		taskUsecase, _ := usecase.NewTask(taskRepository, taskRepository, taskRepository, taskRepository, userRepository, encryptor)
		tasks, _ := taskUsecase.ListByUserID(context.Background(), 3)

		c.JSON(http.StatusOK, tasks)
	}
}

func postTasks(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskRepository, _ := repository.NewTask(db)
		userRepository, _ := repository.NewUser(db)

		encryptor, _ := encryption.New("123456789123456789123456")

		taskUsecase, _ := usecase.NewTask(taskRepository, taskRepository, taskRepository, taskRepository, userRepository, encryptor)
		tasks, _ := taskUsecase.Add(context.Background(), domain.Task{
			Summary: "This is User Task 02",
			UserID:  2,
		})

		c.JSON(http.StatusOK, tasks)
	}
}
