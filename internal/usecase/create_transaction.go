package usecase

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/seuuser/go-event-platform/internal/domain"
	"github.com/seuuser/go-event-platform/internal/port"
)

type CreateTransaction struct {
	repo      port.TransactionRepository
	publisher port.Publisher
}

func NewCreateTransaction(r port.TransactionRepository, p port.Publisher) *CreateTransaction {
	return &CreateTransaction{
		repo:      r,
		publisher: p,
	}
}

func (uc *CreateTransaction) Execute(amount float64) (string, error) {
	id := uuid.New().String()

	t := &domain.Transaction{
		ID:     id,
		Amount: amount,
		Status: domain.Pending,
	}

	if err := uc.repo.Save(t); err != nil {
		return "", err
	}

	body, _ := json.Marshal(t)

	if err := uc.publisher.Publish("transactions", body); err != nil {
		return "", err
	}

	return id, nil

}
