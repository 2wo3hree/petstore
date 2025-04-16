package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"petstore/internal/models"
	"petstore/internal/responder"
	"petstore/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PetHandler struct {
	service   service.PetService
	responder responder.Responder
}

func NewPetHandler(s service.PetService, r responder.Responder) *PetHandler {
	return &PetHandler{
		service:   s,
		responder: r,
	}
}

// Create godoc
// @Summary Add a new pet to the store
// @Tags pet
// @Accept json
// @Produce json
// @Param pet body models.Pet true "Pet object"
// @Success 201 {object} map[string]string
// @Security ApiKeyAuth
// @Router /pet [post]
func (h *PetHandler) Create(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid input"))
		return
	}
	if err := h.service.Create(r.Context(), pet); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusCreated, map[string]string{"message": "pet created"})
}

// Update godoc
// @Summary Update an existing pet
// @Tags pet
// @Accept json
// @Produce json
// @Param pet body models.Pet true "Pet object"
// @Success 200 {object} map[string]string
// @Security ApiKeyAuth
// @Router /pet [put]
func (h *PetHandler) Update(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid input"))
		return
	}
	if err := h.service.Update(r.Context(), pet); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, map[string]string{"message": "pet updated"})
}

// UpdateByID godoc
// @Summary Update pet by ID
// @Tags pet
// @Accept json
// @Produce json
// @Param petId path int true "Pet ID"
// @Param pet body models.Pet true "Updated pet object"
// @Success 200 {object} map[string]string
// @Security ApiKeyAuth
// @Router /pet/{petId} [post]
func (h *PetHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid input"))
		return
	}
	if err := h.service.Update(r.Context(), pet); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, map[string]string{"message": "pet updated by ID"})
}

// GetByID godoc
// @Summary Find pet by ID
// @Tags pet
// @Produce json
// @Param petId path int true "Pet ID"
// @Success 200 {object} models.Pet
// @Security ApiKeyAuth
// @Router /pet/{petId} [get]
func (h *PetHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "petId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid ID"))
		return
	}

	pet, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.responder.Error(w, http.StatusNotFound, fmt.Errorf("pet not found"))
		return
	}
	h.responder.JSON(w, http.StatusOK, pet)
}

// Delete godoc
// @Summary Deletes a pet
// @Tags pet
// @Param petId path int true "Pet ID"
// @Success 204 {string} string "No Content"
// @Security ApiKeyAuth
// @Router /pet/{petId} [delete]
func (h *PetHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "petId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("invalid ID"))
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusNoContent, nil)
}

// FindByStatus godoc
// @Summary Finds pets by status
// @Tags pet
// @Produce json
// @Param status query string true "Status value"
// @Success 200 {array} models.Pet
// @Security ApiKeyAuth
// @Router /pet/findByStatus [get]
func (h *PetHandler) FindByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		h.responder.Error(w, http.StatusBadRequest, fmt.Errorf("status is required"))
		return
	}

	pets, err := h.service.FindByStatus(r.Context(), status)
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, pets)
}
