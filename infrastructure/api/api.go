package api

import (
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateTaskRoutes(r *gin.Engine) {
	tasks := r.Group("tasks")

	tasks.GET("", getTasks())
}

func getTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentTime := time.Now()
		c.JSON(http.StatusOK, domain.Task{
			ID:      1,
			Summary: "Dummy Summary",
			Date:    &currentTime,
			UserID:  1,
		})
	}
}
