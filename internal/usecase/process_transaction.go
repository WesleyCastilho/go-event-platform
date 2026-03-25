package usecase

import (
	"encoding/json"
	"time"

	"github.com/seuuser/go-event-platform/internal/domain"
	"github.com/seuuser/go-event-platform/internal/port"
)

type ProcessTransaction struct {
	repo port.TransactionRepository
}

func NewProcessTransaction(r port.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{repo: r}
}

// Lembrem us é o Use Case
func (uc *ProcessTransaction) Execute(body []byte) error {
	var t domain.Transaction
	json.Unmarshal(body, &t)

	//Simular o processamento
	/// tem que por algo mais legal aqui ... esse delay so nao fica legal nao...
	time.Sleep(500 * time.Millisecond)

	return uc.repo.UpdateStatus(t.ID, domain.Processed)
}
