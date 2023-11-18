package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SupAdminAuth(c *gin.Context) {
	tokenString, err := c.Cookie("supadminToken")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	supadminId, err := ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("supadminId", supadminId)
	c.Next()
}