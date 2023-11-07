package handlerutil

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("userId")
	adminID, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return adminID, err
}