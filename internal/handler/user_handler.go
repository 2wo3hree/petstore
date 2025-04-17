package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"petstore/internal/models"
	"petstore/internal/responder"
	"petstore/internal/service"
)

type UserHandler struct {
	service   service.UserService
	responder responder.Responder
}

func NewUserHandler(s service.UserService, r responder.Responder) *UserHandler {
	return &UserHandler{service: s, responder: r}
}

// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User name only"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal error"
// @Router /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.responder.Error(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.Create(r.Context(), models.User{Name: req.Name})
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusCreated, map[string]interface{}{"id": id, "message": "user created"})
}

// @Summary List users
// @Tags users
// @Produce json
// @Param limit query int false "Max items"
// @Param offset query int false "Offset"
// @Success 200 {object} models.ListUsersResponse
// @Failure 500 {string} string "internal error"
// @Router /users [get]
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	limit, offset := 10, 0
	if s := r.URL.Query().Get("limit"); s != "" {
		if v, err := strconv.Atoi(s); err == nil {
			limit = v
		}
	}
	if s := r.URL.Query().Get("offset"); s != "" {
		if v, err := strconv.Atoi(s); err == nil {
			offset = v
		}
	}

	users, total, err := h.service.List(r.Context(), limit, offset)
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}

	// fill rentals for each
	for i := range users {
		u, err := h.service.GetWithRentals(r.Context(), users[i].ID)
		if err != nil {
			h.responder.Error(w, http.StatusInternalServerError, err)
			return
		}
		users[i].RentedBooks = u.RentedBooks
	}

	h.responder.JSON(w, http.StatusOK, models.ListUsersResponse{Total: total, Users: users})
}

// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.responder.Error(w, http.StatusBadRequest, err)
		return
	}
	u, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.responder.Error(w, http.StatusNotFound, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, u)
}

// @Summary Get user with their rented books
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "not found"
// @Router /users/{id}/rentals [get]
func (h *UserHandler) GetWithRentals(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.responder.Error(w, http.StatusBadRequest, err)
		return
	}
	u, err := h.service.GetWithRentals(r.Context(), id)
	if err != nil {
		h.responder.Error(w, http.StatusNotFound, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, u)
}
