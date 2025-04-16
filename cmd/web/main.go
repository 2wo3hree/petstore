// @title PetStore API
// @version 1.0
// @description This is a sample API server for petStores.
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	_ "petstore/docs"
	"petstore/internal/app"
	"petstore/internal/config"
	"petstore/internal/server"
	"petstore/internal/shutdown"
)

func main() {
	cfg := config.Load()

	application := app.NewApp(cfg)

	s := server.NewServer(":8080", application.Router)

	shutdown.Gracefully(s)
}
