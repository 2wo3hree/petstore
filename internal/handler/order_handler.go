package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"petstore/internal/responder"
	"strconv"

	"github.com/go-chi/chi/v5"
	"petstore/internal/models"
	"petstore/internal/service"
)

type OrderHandler struct {
	service   service.OrderService
	responder responder.Responder
}

func NewOrderHandler(s service.OrderService, r responder.Responder) *OrderHandler {
	return &OrderHandler{
		service:   s,
		responder: r,
	}
}

// Create godoc
// @Summary Create order
// @Tags store
// @Accept json
// @Produce json
// @Param order body models.Order true "Order object"
// @Success 200 {string} string "ok"
// @Router /store/order [post]
func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid input"))
		return
	}
	if err := h.service.Create(r.Context(), order); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("failed to create user"))
		return
	}
	h.responder.JSON(w, http.StatusOK, nil)
}

// GetByID godoc
// @Summary Get order by ID
// @Tags store
// @Produce json
// @Param orderId path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /store/order/{orderId} [get]
func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	order, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.responder.Error(w, http.StatusNotFound, fmt.Errorf("order not found"))
		return
	}
	h.responder.JSON(w, http.StatusOK, order)
}

// Delete godoc
// @Summary Delete order
// @Tags store
// @Param orderId path int true "Order ID"
// @Success 200 {string} string "deleted"
// @Router /store/order/{orderId} [delete]
func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := h.service.Delete(r.Context(), id); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("delete error"))
		return
	}
	h.responder.JSON(w, http.StatusNoContent, nil)
}

// GetInventory godoc
// @Summary Get inventory
// @Tags store
// @Produce json
// @Success 200 {object} map[string]int
// @Security ApiKeyAuth
// @Router /store/inventory [get]
func (h *OrderHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	inventory, err := h.service.GetInventory(ctx)
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, fmt.Errorf("Ошибка получения инвентаря"))
		return
	}

	h.responder.JSON(w, http.StatusOK, inventory)
}
