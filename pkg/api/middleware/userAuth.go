package middleware

import (
	"fmt"
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

func TestUserAuth(c *gin.Context) {
	c.Set("userId", 1)
	fmt.Println("calll")
	c.Next()
}