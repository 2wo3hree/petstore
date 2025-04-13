// @title User Management API
// @version 1.0
// @description This is a sample API server for managing users.
// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	_ "repo/docs"
	"repo/internal/app"
	"repo/internal/config"
	"repo/internal/server"
	"repo/internal/shutdown"
)

func main() {
	cfg := config.Load()

	application := app.NewApp(cfg)

	s := server.NewServer(":8080", application.Router)

	shutdown.Gracefully(s)
}
