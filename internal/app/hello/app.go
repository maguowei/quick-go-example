package hello

import (
	"github.com/gin-gonic/gin"
	"example/internal/pkg/configs"
	"example/internal/app/hello/routers"
)

var app App

type App struct {
	appName string
	addr   string
	router *gin.Engine
}

func (app *App) init() {
	routers.InitRouter()
	configs.InitConfig()
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
