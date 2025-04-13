package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"repo/internal/model"
	"repo/internal/responder"
	"repo/internal/service"
	"strconv"
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

// @Summary Create user
// @Description Создать нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User data"
// @Success 201 {object} map[string]string "message"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users [post]
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		u.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid input"))
		return
	}

	if err := u.service.Create(r.Context(), user); err != nil {
		u.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("failed to create user"))
		return
	}
	u.responder.JSON(w, http.StatusOK, map[string]string{"message": "User create"})
}

// @Summary Get user by ID
// @Description Получить пользователя по ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users/{id} [get]
func (u *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := u.service.GetByID(r.Context(), id)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	u.responder.JSON(w, http.StatusOK, user)
}

// @Summary Update user
// @Description Обновить имя и email пользователя по ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body model.User true "User data to update"
// @Success 200 {object} map[string]string "message"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users/{id} [put]
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		u.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid input"))
		return
	}
	user.ID = id

	err := u.service.Update(r.Context(), user)
	if err != nil {
		u.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("update error"))
		return
	}

	u.responder.JSON(w, http.StatusOK, map[string]string{"message": "User updated"})
}

// @Summary Delete user
// @Description Мягкое удаление пользователя (проставляется deleted_at)
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string "message"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users/{id} [delete]
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := u.service.Delete(r.Context(), id)
	if err != nil {
		u.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("delete error"))
		return
	}

	u.responder.JSON(w, http.StatusOK, map[string]string{"message": "User deleted"})
}

// @Summary List users
// @Description возвращает список всех пользователей, с указанием количества пользователей в ответе и пагинацией (параметры limit и offset)
// @Tags users
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} model.ListUsersResponse
// @Failure 500 {string} string "Internal error"
// @Router /api/users [get]
func (u *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10
	offset := 0

	if limitStr != "" {
		if parsed, err := strconv.Atoi(limitStr); err == nil {
			limit = parsed
		}
	}
	if offsetStr != "" {
		if parsed, err := strconv.Atoi(offsetStr); err == nil {
			offset = parsed
		}
	}

	users, count, err := u.service.List(r.Context(), limit, offset)
	if err != nil {
		u.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch users"))
		return
	}

	resp := model.ListUsersResponse{
		Total: count,
		Users: users,
	}

	u.responder.JSON(w, http.StatusOK, resp)
}
