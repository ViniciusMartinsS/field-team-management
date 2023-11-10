package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTaskRoutes(r *gin.Engine) {
	tasks := r.Group("tasks")

	tasks.GET("", getTasks())
}

func getTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "tasks")
	}
}
