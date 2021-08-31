package hello

import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"example/internal/app/hello/api"
)

var app App

type App struct {
	appName string
	addr   string
	router *gin.Engine
}

func (app *App) init() {
	app.router = gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(app.router)

	app.router.Use(gin.Logger(), gin.Recovery())

	app.router.GET("/", api.Index)
	app.router.GET("/health", api.Health)
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
