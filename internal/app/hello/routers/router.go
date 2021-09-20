package routers

import (
	"example/internal/app/hello/restapi"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var router *gin.Engine


func InitRouter() {
	router = gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", restapi.Index)
	router.GET("/health", restapi.Health)
}

func GetRouter() *gin.Engine {
	return router
}
