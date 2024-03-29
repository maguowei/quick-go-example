package example

import (
	"github.com/gin-gonic/gin"
	"example/internal/example/interfaces/server"
	"example/internal/pkg/app"
	"example/internal/pkg/configs"
)


type defaultApp struct {
	appName string
	addr    string
	server  *gin.Engine
}

func (app *defaultApp) Run() error {
	server.InitServer()
	app.server = server.GetServer()
	if err := app.server.Run(); err != nil {
		return err
	}
	return nil
}

func NewApp() app.App {
	configs.InitConfig()
	myApp := &defaultApp{}
	return myApp
}
