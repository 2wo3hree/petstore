// internal/handler/author_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"petstore/internal/models"
	"petstore/internal/responder"
	"petstore/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AuthorHandler struct {
	service   service.Facade
	responder responder.Responder
}

func NewAuthorHandler(s service.Facade, r responder.Responder) *AuthorHandler {
	return &AuthorHandler{service: s, responder: r}
}

// CreateAuthor godoc
// @Summary   Create a new author
// @Tags      authors
// @Accept    json
// @Produce   json
// @Param     author  body     models.CreateAuthorRequest true "Only author name"
// @Success   201     {object} map[string]interface{}
// @Failure   400     {object} map[string]string
// @Failure   500     {object} map[string]string
// @Router    /authors [post]
func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.responder.Error(w, http.StatusBadRequest, err)
		return
	}
	a := models.Author{Name: req.Name}
	id, err := h.service.CreateAuthor(r.Context(), a)
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusCreated, map[string]interface{}{"id": id, "message": "user created"})
}

// ListAuthors godoc
// @Summary   List authors
// @Tags      authors
// @Accept    json
// @Produce   json
// @Param     limit  query int false "Max items"
// @Param     offset query int false "Offset"
// @Success   200     {object} []models.Author
// @Failure   500     {object} map[string]string
// @Router    /authors [get]
func (h *AuthorHandler) ListAuthors(w http.ResponseWriter, r *http.Request) {
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
	list, _, err := h.service.ListAuthors(r.Context(), limit, offset)
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, list)
}

// GetAuthor godoc
// @Summary   Get author by ID
// @Tags      authors
// @Accept    json
// @Produce   json
// @Param     id   path    int  true  "Author ID"
// @Success   200  {object} models.Author
// @Failure   400  {object} map[string]string
// @Failure   404  {object} map[string]string
// @Router    /authors/{id} [get]
func (h *AuthorHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		h.responder.Error(w, http.StatusBadRequest, err)
		return
	}
	a, err := h.service.GetAuthorByID(r.Context(), id)
	if err != nil {
		h.responder.Error(w, http.StatusNotFound, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, a)
}

// @Summary   List top authors by rentals
// @Tags      library
// @Accept    json
// @Produce   json
// @Success   200 {array} models.AuthorCount
// @Failure   500 {object} map[string]string
// @Router    /library/top [get]
func (h *AuthorHandler) TopAuthors(w http.ResponseWriter, r *http.Request) {
	top, err := h.service.GetTopAuthors(r.Context())
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, top)
}
