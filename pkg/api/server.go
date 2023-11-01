package http

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	userHandler *handler.UserHandler) *ServerHTTP {

	engine := gin.Default()

	return &ServerHTTP{engine: engine}
}
func (sh *ServerHTTP) Start() {
	sh.engine.Run()
}
