package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"petstore/internal/config"
	"petstore/internal/db"
	"petstore/internal/facade"
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

	db.SeedDatabase(pool)

	// init repo
	userRepo := postgres.NewUserRepo(pool)
	authorRepo := postgres.NewAuthorRepo(pool)
	bookRepo := postgres.NewBookRepo(pool)
	rentalRepo := postgres.NewRentalRepo(pool)
	// init service
	userService := service.NewUserService(userRepo)
	authorService := service.NewAuthorService(authorRepo)
	bookService := service.NewBookService(bookRepo)
	rentalService := service.NewRentalService(rentalRepo, userRepo, bookRepo)

	// init super service (суперсервис)
	libraryService := service.NewLibraryService(userService, bookService, rentalService, authorService)

	// init facade (фасад работает поверх суперсервиса)
	libraryFacade := facade.NewFacade(libraryService)

	// init responder
	resp := responder.NewJSONResponder()

	// init handlers
	userHandler := handler.NewUserHandler(userService, resp)
	authorHandler := handler.NewAuthorHandler(authorService, resp)
	bookHandler := handler.NewBookHandler(bookService, resp)

	facadeHandler := handler.NewFacadeHandler(libraryFacade, resp)

	//auth.InitJWT()

	// init router
	r := router.SetupRouter(userHandler, authorHandler, bookHandler, facadeHandler)

	return &App{
		Router: r,
		//JWTAuth: auth.TokenAuth,
	}
}
