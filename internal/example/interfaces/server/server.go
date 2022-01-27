package server

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maguowei/example/internal/example/application/service"
	"github.com/maguowei/example/internal/example/domain/repository"
	"github.com/maguowei/example/internal/example/domain/repository/ent"
	domainService "github.com/maguowei/example/internal/example/domain/service"
	"github.com/maguowei/example/internal/example/interfaces/restapi"
	"github.com/spf13/viper"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"log"
)

var server *gin.Engine


func InitServer() {
	server = gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(server)

	server.Use(gin.Logger(), gin.Recovery())

	client, err := ent.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
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
