package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"petstore/internal/config"
	"petstore/internal/db"
	"petstore/internal/handler"
	"petstore/internal/repository/postgres"
	"petstore/internal/responder"
	"petstore/internal/router"
	"petstore/internal/service"
)

type App struct {
	Router *chi.Mux
	//JWTAuth *jwtauth.JWTAuth
}

func NewApp(cfg *config.Config) *App {
	// int db
	pool := db.NewPostgres(cfg)

	// run Migrations
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db.RunMigrations(dbURL)

	// init repo
	userRepo := postgres.NewUserRepo(pool)
	petRepo := postgres.NewPetRepo(pool)
	orderRepo := postgres.NewOrderRepo(pool)

	// init service
	userService := service.NewUserService(userRepo)
	petService := service.NewPetService(petRepo)
	orderService := service.NewOrderService(orderRepo)

	// init responder
	resp := responder.NewJSONResponder()

	// init handlers
	userHandler := handler.NewUserHandler(userService, resp)
	petHandler := handler.NewPetHandler(petService, resp)
	orderHandler := handler.NewOrderHandler(orderService, resp)

	//auth.InitJWT()

	// init router
	r := router.SetupRouter(petHandler, userHandler, orderHandler, userService)

	return &App{
		Router: r,
		//JWTAuth: auth.TokenAuth,
	}
}
