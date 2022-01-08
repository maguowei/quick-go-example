package example

import (
	"github.com/gin-gonic/gin"
	"github.com/maguowei/example/internal/pkg/app"
	"github.com/maguowei/example/internal/example/interfaces/routers"
	"github.com/maguowei/example/internal/pkg/configs"
)

var myApp app.App

type DefaultApp struct {
	appName string
	addr   string
	router *gin.Engine
}

func (app *DefaultApp) init() {
	configs.InitConfig()
	routers.InitRouter()
	app.router = routers.GetRouter()
}

func (app *DefaultApp) Run() {
	app.init()
	app.router.Run()
}

func NewApp() app.App {
	myApp = &DefaultApp{}
	return myApp
}

func GetApp() app.App {
	return myApp
}
