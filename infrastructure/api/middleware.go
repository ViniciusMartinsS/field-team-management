package api

import (
	jwtAuthenticator "github.com/ViniciusMartinss/field-team-management/infrastructure/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, "")
			c.Abort()
			return
		}

		if strings.Contains(token, "Bearer") {
			token = strings.TrimSpace(strings.Replace(token, "Bearer", "", -1))
		}

		authenticator, err := jwtAuthenticator.New("my_secret_key")
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			c.Abort()
			return
		}

		valid, claims, err := authenticator.IsAccessTokenValid(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			c.Abort()
			return
		}

		if !valid {
			c.JSON(http.StatusInternalServerError, "")
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("role_id", claims["role_id"])

		c.Next()
	}
}
