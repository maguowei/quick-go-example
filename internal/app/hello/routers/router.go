package routers

import (
	"example/internal/app/hello/api"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var router *gin.Engine


func InitRouter() {
	router = gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", api.Index)
	router.GET("/health", api.Health)
}

func GetRouter() *gin.Engine {
	return router
}
