package api

import (
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const unauthorizedMessage = "unauthorized"

func Authenticator(authenticator domain.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, toResponse(false, unauthorizedMessage))
			c.Abort()
			return
		}

		if strings.Contains(token, "Bearer") {
			token = strings.TrimSpace(strings.Replace(token, "Bearer", "", -1))
		}

		valid, claims, err := authenticator.IsAccessTokenValid(token)
		if !valid || err != nil {
			c.JSON(http.StatusUnauthorized, toResponse(false, unauthorizedMessage))
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("role_id", claims["role_id"])

		c.Next()
	}
}
