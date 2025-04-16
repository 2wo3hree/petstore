package auth

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"petstore/internal/service"
	_ "petstore/internal/service"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthHandler struct {
	TokenAuth   *jwtauth.JWTAuth
	UserService service.UserService
}

// LoginHandler godoc
// @Summary Login user
// @Tags user
// @Produce json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "bad request"
// @Router /user/login [get]
func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username == "" || password == "" {
		http.Error(w, "не хватает username или password", http.StatusBadRequest)
		return
	}

	user, err := a.UserService.GetByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, "пользователь не найден", http.StatusUnauthorized)
		return
	}

	// Сравнение паролей
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		http.Error(w, "неверный пароль", http.StatusUnauthorized)
		return
	}

	if a.TokenAuth == nil {
		http.Error(w, "TokenAuth не инициализирован", http.StatusInternalServerError)
		return
	}

	// Генерация токена
	_, token, _ := a.TokenAuth.Encode(map[string]interface{}{"sub": username})

	// Установка токена в куку
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "logged in"})
}

// LogoutHandler godoc
// @Summary Logout user
// @Tags user
// @Success 200 {object} map[string]string
// @Router /user/logout [get]
func (a *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Сбросить cookie с токеном
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "logged out"})
}
