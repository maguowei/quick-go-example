package example

import (
	"github.com/gin-gonic/gin"
	"github.com/maguowei/example/internal/pkg/app"
	"github.com/maguowei/example/internal/example/interfaces/server"
	"github.com/maguowei/example/internal/pkg/configs"
)

var myApp app.App

type DefaultApp struct {
	appName string
	addr   string
	server *gin.Engine
}

func (app *DefaultApp) init() {
	configs.InitConfig()
	server.InitServer()
	app.server = server.GetServer()
}

func (app *DefaultApp) Run() {
	app.init()
	app.server.Run()
}

func NewApp() app.App {
	myApp = &DefaultApp{}
	return myApp
}

func GetApp() app.App {
	return myApp
}
