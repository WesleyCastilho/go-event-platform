package http

import (
	"encoding/json"
	"net/http"

	"github.com/seuuser/go-event-platform/internal/usecase"
)

type Handler struct {
	createUC *usecase.CreateTransaction
}

func NewHandler(uc *usecase.CreateTransaction) *Handler {
	return &Handler{createUC: uc}
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Amount float64 `json:"amount"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	id, err := h.createUC.Execute(input.Amount)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(id))
}
