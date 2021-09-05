package hello

import (
	"example/internal/app/hello/routers"
	"example/internal/pkg/configs"
	"github.com/gin-gonic/gin"
)

var app App

type App struct {
	appName string
	addr   string
	router *gin.Engine
}

func (app *App) init() {
	configs.InitConfig()
	routers.InitRouter()
	app.router = routers.GetRouter()
}

func (app *App) Run() {
	app.init()
	app.router.Run()
}

func NewApp() *App {
	app = App{}
	return &app
}

func GetApp() *App {
	return &app
}
