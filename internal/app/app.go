package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"repo/internal/config"
	"repo/internal/db"
	"repo/internal/handler"
	"repo/internal/repository/postgres"
	"repo/internal/responder"
	"repo/internal/router"
	"repo/internal/service"
)

type App struct {
	Router   *chi.Mux
	Handlers *handler.UserHandler
}

func NewApp(cfg *config.Config) *App {
	// int db
	pool := db.NewPostgres(cfg)

	// run Migrations
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db.RunMigrations(dbURL)

	// init repo
	userRepo := postgres.NewUserRepository(pool)

	// init service
	s := service.NewUserService(userRepo)

	// init responder
	resp := responder.NewJSONResponder()

	// init handlers
	h := handler.NewUserHandler(s, resp)

	// init router
	r := router.SetupRouter(h)

	return &App{
		Router:   r,
		Handlers: h,
	}
}
