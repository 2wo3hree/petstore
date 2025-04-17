package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"petstore/internal/models"
	"petstore/internal/responder"
	"petstore/internal/service"
	"strconv"
)

type BookHandler struct {
	service   service.BookService
	responder responder.Responder
}

func NewBookHandler(s service.BookService, r responder.Responder) *BookHandler {
	return &BookHandler{
		service:   s,
		responder: r,
	}
}

// CreateBook godoc
// @Summary Add a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.CreateBookRequest true "Book payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal error"
// @Router /books [post]
func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.responder.Error(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.Create(r.Context(), models.Book{
		Title:    req.Title,
		AuthorID: req.AuthorID,
	})
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusCreated, map[string]interface{}{
		"id":      id,
		"message": "book created",
	})
}

// ListBooks godoc
// @Summary List all books
// @Tags books
// @Produce json
// @Param limit  query int false "Max items"
// @Param offset query int false "Offset"
// @Success 200 {object} models.ListBooksResponse
// @Failure 500 {string} string "Internal error"
// @Router /books [get]
func (h *BookHandler) List(w http.ResponseWriter, r *http.Request) {
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

	books, total, err := h.service.List(r.Context(), limit, offset)
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	resp := models.ListBooksResponse{Total: total, Books: books}
	h.responder.JSON(w, http.StatusOK, resp)
}

// GetBookByID godoc
// @Summary Get book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not found"
// @Router /books/{id} [get]
func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.responder.Error(w, http.StatusBadRequest, err)
		return
	}
	book, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.responder.Error(w, http.StatusNotFound, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, book)
}
