package server

import (
	"github.com/gin-gonic/gin"
	"github.com/maguowei/example/internal/example/interfaces/restapi"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var server *gin.Engine


func InitServer() {
	server = gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(server)

	server.Use(gin.Logger(), gin.Recovery())

	server.GET("/", restapi.Index)
	server.GET("/health", restapi.Health)
}

func GetServer() *gin.Engine {
	return server
}
