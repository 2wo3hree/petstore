package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/swaggo/http-swagger"
	"net/http"
	"petstore/internal/auth"
	"petstore/internal/handler"
	"petstore/internal/service"
)

// TokenFromCookie — middleware, который достаёт токен из cookie и добавляет его в Authorization header
func TokenFromCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err == nil {
			r.Header.Set("Authorization", "Bearer "+cookie.Value)
		}
		next.ServeHTTP(w, r)
	})
}

func SetupRouter(pet *handler.PetHandler, user *handler.UserHandler, order *handler.OrderHandler, userService service.UserService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	auth.InitJWT()

	authHandler := &auth.AuthHandler{
		TokenAuth:   auth.TokenAuth,
		UserService: userService,
	}

	//Auth
	r.Get("/user/login", authHandler.LoginHandler)
	r.Get("/user/logout", authHandler.LogoutHandler)

	//Protected routes
	r.Group(func(r chi.Router) {
		r.Use(TokenFromCookie)
		r.Use(jwtauth.Verifier(authHandler.TokenAuth))
		r.Use(jwtauth.Authenticator)

		// Pet routes
		r.Route("/pet", func(r chi.Router) {
			r.Post("/", pet.Create)
			r.Put("/", pet.Update)
			r.Get("/findByStatus", pet.FindByStatus)
			r.Get("/{petId}", pet.GetByID)
			r.Post("/{petId}", pet.UpdateByID)
			r.Delete("/{petId}", pet.Delete)
		})

		r.Get("/store/inventory", order.GetInventory)

	})

	// Store routes
	r.Route("/store", func(r chi.Router) {
		r.Post("/order", order.Create)
		r.Get("/order/{orderId}", order.GetByID)
		r.Delete("/order/{orderId}", order.Delete)
	})

	// User routes
	r.Route("/user", func(r chi.Router) {
		r.Post("/", user.Create)
		r.Post("/createWithArray", user.CreateWithArray)
		r.Post("/createWithList", user.CreateWithList)
		//r.Get("/logout", user.Logout)
		r.Get("/{username}", user.GetByUsername)
		r.Put("/{username}", user.Update)
		r.Delete("/{username}", user.Delete)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
