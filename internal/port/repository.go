package port

import "github.com/seuuser/go-event-platform/internal/domain"

type TransactionRepository interface {
	Save(t *domain.Transaction) error
	UpdateStatus(id string, status domain.Status) error
}
