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
}

func NewApp(cfg *config.Config) *App {
	// int db
	pool := db.NewPostgres(cfg)

	// run Migrations
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db.RunMigrations(dbURL)

	db.SeedDatabase(pool)

	// init repo
	userRepo := postgres.NewUserRepo(pool)
	authorRepo := postgres.NewAuthorRepo(pool)
	bookRepo := postgres.NewBookRepo(pool)

	// init service
	userService := service.NewUserService(userRepo)
	authorService := service.NewAuthorService(authorRepo)
	bookService := service.NewBookService(bookRepo)

	// init super service (суперсервис)
	libraryService := service.NewLibrarySuperService(userService, bookService, authorService)

	// init responder
	resp := responder.NewJSONResponder()

	// init handlers
	userHandler := handler.NewUserHandler(libraryService, resp)
	authorHandler := handler.NewAuthorHandler(libraryService, resp)
	bookHandler := handler.NewBookHandler(libraryService, resp)

	// init router
	r := router.SetupRouter(userHandler, authorHandler, bookHandler)

	return &App{
		Router: r,
	}
}
