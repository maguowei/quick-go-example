package server

import (
    "os"
	"context"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"example/internal/example/application/service"
	domainService "example/internal/example/domain/service"
	"example/internal/example/infrastructure/persistence/ent"
	"example/internal/example/infrastructure/log"
	"example/internal/example/infrastructure/repository"
	"example/internal/example/interfaces/restapi"
	"github.com/spf13/viper"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var server *gin.Engine


func InitServer() {
	server = gin.New()
	server.Use(requestid.New())
	p := ginprometheus.NewPrometheus("gin")
	p.Use(server)

	server.Use(gin.Logger(), gin.Recovery())

	appName := os.Getenv("APP_NAME")
	appEnv := os.Getenv("APP_ENV")

	if appName == "" {
	    appName = "example"
	}

	if appEnv == "" {
	    appEnv = "local"
	}

	logger := log.NewLogger(appName, appEnv, ".")
	logger.Info("Starting server...")

	client, err := ent.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {
		logger.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatalf("failed creating schema resources: %v", err)
	}

	exampleRepository := repository.NewExampleRepository(client)
	exampleDomainService := domainService.NewExampleDomainService(exampleRepository)
	exampleAppService := service.NewExampleAppService(exampleDomainService)
	exampleApi := restapi.NewExampleApi(exampleAppService)


	server.GET("/", restapi.Index)
	server.GET("/health", restapi.Health)
	server.POST("/example/create", exampleApi.CreateExample)
}

func GetServer() *gin.Engine {
	return server
}
