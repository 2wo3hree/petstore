package responder

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	JSON(w http.ResponseWriter, status int, data interface{})
	Error(w http.ResponseWriter, status int, err error)
}

type JSONResponder struct{}

func NewJSONResponder() Responder {
	return &JSONResponder{}
}

func (jr *JSONResponder) JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (jr *JSONResponder) Error(w http.ResponseWriter, status int, err error) {
	jr.JSON(w, status, map[string]string{"error": err.Error()})
}
