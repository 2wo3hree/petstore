package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"petstore/internal/models"
	"petstore/internal/responder"
	"petstore/internal/service"
)

type UserHandler struct {
	service   service.UserService
	responder responder.Responder
}

func NewUserHandler(s service.UserService, r responder.Responder) *UserHandler {
	return &UserHandler{
		service:   s,
		responder: r,
	}
}

// Create godoc
// @Summary Create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} map[string]string
// @Router /user [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid input"))
		return
	}
	if err := h.service.Create(r.Context(), user); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}

// CreateWithArray godoc
// @Summary Create user with array
// @Tags user
// @Accept json
// @Produce json
// @Param users body []models.User true "Array of user objects"
// @Success 201 {object} map[string]string
// @Router /user/createWithArray [post]
func (h *UserHandler) CreateWithArray(w http.ResponseWriter, r *http.Request) {
	h.Create(w, r) // логика та же
}

// CreateWithList godoc
// @Summary Create user with list
// @Tags user
// @Accept json
// @Produce json
// @Param users body []models.User true "List of user objects"
// @Success 201 {object} map[string]string
// @Router /user/createWithList [post]
func (h *UserHandler) CreateWithList(w http.ResponseWriter, r *http.Request) {
	h.Create(w, r) // логика та же
}

// GetByUsername godoc
// @Summary Get user by username
// @Tags user
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.User
// @Router /user/{username} [get]
func (h *UserHandler) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := h.service.GetByUsername(r.Context(), username)
	if err != nil {
		h.responder.Error(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}
	h.responder.JSON(w, http.StatusOK, user)
}

// Update godoc
// @Summary Update user
// @Tags user
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param user body models.User true "Updated user object"
// @Success 200 {object} map[string]string
// @Router /user/{username} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid input"))
		return
	}

	if err := h.service.Update(r.Context(), username, updated); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("failed to update user"))
		return
	}

	h.responder.JSON(w, http.StatusOK, map[string]string{"message": "updated"})
}

// Delete godoc
// @Summary Delete user
// @Tags user
// @Param username path string true "Username"
// @Success 204 {string} string "No Content"
// @Router /user/{username} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if err := h.service.Delete(r.Context(), username); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusNoContent, nil)
}

// Logout — пока временно не используется
/*
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Logout(r.Context()); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("logout error"))
		return
	}
	h.responder.JSON(w, http.StatusOK, map[string]string{"message": "logged out"})
}
*/
