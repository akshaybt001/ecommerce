package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context){
	tokenString, err :=c.Cookie("userToken")
	if err !=nil{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId,err:= ValidateToken(tokenString)
	if err !=nil{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("userId",userId)
	c.Next()
}