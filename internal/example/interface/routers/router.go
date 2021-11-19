package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maguowei/example/internal/example/interface/restapi"
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
