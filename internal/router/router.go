package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger"
	"petstore/internal/handler"
)

//// TokenFromCookie — middleware, который достаёт токен из cookie и добавляет его в Authorization header
//func TokenFromCookie(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		cookie, err := r.Cookie("jwt")
//		if err == nil {
//			r.Header.Set("Authorization", "Bearer "+cookie.Value)
//		}
//		next.ServeHTTP(w, r)
//	})
//}

func SetupRouter(userHandler *handler.UserHandler, authorHandler *handler.AuthorHandler, bookHandler *handler.BookHandler, facadeHandler *handler.FacadeHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//auth.InitJWT()
	//
	//authHandler := &auth.AuthHandler{
	//	TokenAuth:   auth.TokenAuth,
	//	UserService: userService,
	//}

	//Auth
	//r.Get("/user/login", authHandler.LoginHandler)
	//r.Get("/user/logout", authHandler.LogoutHandler)

	//Protected routes
	//r.Group(func(r chi.Router) {
	//	r.Use(TokenFromCookie)
	//	r.Use(jwtauth.Verifier(authHandler.TokenAuth))
	//	r.Use(jwtauth.Authenticator)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Get("/", userHandler.List)
		r.Get("/{id}", userHandler.GetByID)
		r.Get("/{id}/rentals", userHandler.GetWithRentals)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Post("/", authorHandler.CreateAuthor)
		r.Get("/", authorHandler.ListAuthors)
		r.Get("/{id}", authorHandler.GetAuthor)
	})

	r.Route("/books", func(r chi.Router) {
		r.Post("/", bookHandler.Create)
		r.Get("/", bookHandler.List)
		r.Get("/{id}", bookHandler.GetByID)
	})

	r.Route("/library", func(r chi.Router) {
		r.Post("/issue/{userId}/{bookId}", facadeHandler.IssueBook)
		r.Post("/return/{userId}/{bookId}", facadeHandler.ReturnBook)
		r.Get("/top", facadeHandler.TopAuthors)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
