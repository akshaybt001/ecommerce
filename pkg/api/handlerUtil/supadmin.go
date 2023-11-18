package handlerutil

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSupAdminIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("supadminId")
	adminID, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return adminID, err
}
