package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/keploy/go-sdk/integrations/kgin/v1"
	"github.com/keploy/go-sdk/keploy"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//connecting keploy
	k := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "gin-mongo-api",
			Port: "8080",
		},
		Server: keploy.ServerConfig{
			URL: "http://localhost:8081/api",
		},
	})
	kgin.GinV1(k, router)

	//routes
	routes.UserRoute(router)

	router.Run()
}
