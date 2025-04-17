package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"petstore/internal/facade"
	_ "petstore/internal/models"
	"petstore/internal/responder"
	"strconv"
)

type FacadeHandler struct {
	facade    *facade.Facade
	responder responder.Responder
}

func NewFacadeHandler(f *facade.Facade, r responder.Responder) *FacadeHandler {
	return &FacadeHandler{facade: f, responder: r}
}

// @Summary Issue a book to a user
// @Tags library
// @Produce json
// @Param userId path int true "User ID"
// @Param bookId path int true "Book ID"
// @Success 201 {object} map[string]string
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "User or Book not found"
// @Failure 409 {string} string "Book already issued"
// @Router /library/issue/{userId}/{bookId} [post]
func (h *FacadeHandler) IssueBook(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	bookID, _ := strconv.Atoi(chi.URLParam(r, "bookId"))

	rent, err := h.facade.IssueBook(r.Context(), userID, bookID)
	if err != nil {
		h.responder.Error(w, http.StatusConflict, err)
		return
	}
	h.responder.JSON(w, http.StatusCreated, rent)
}

// @Summary Return a book from a user
// @Tags library
// @Produce json
// @Param userId path int true "User ID"
// @Param bookId path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "No active rental"
// @Router /library/return/{userId}/{bookId} [post]
func (h *FacadeHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	bookID, _ := strconv.Atoi(chi.URLParam(r, "bookId"))

	if err := h.facade.ReturnBook(r.Context(), userID, bookID); err != nil {
		h.responder.Error(w, http.StatusNotFound, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, map[string]string{"message": "returned"})
}

// @Summary   List top authors by rentals
// @Tags      library
// @Accept    json
// @Produce   json
// @Success   200 {array} models.AuthorCount
// @Failure   500 {object} map[string]string
// @Router    /library/top [get]
func (h *FacadeHandler) TopAuthors(w http.ResponseWriter, r *http.Request) {
	top, err := h.facade.TopAuthors(r.Context())
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusOK, top)
}
