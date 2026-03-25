package domain

import "time"

type Status string

const (
	Pending   Status = "PENDING"
	Processed Status = "PROCESSED"
)

type Transaction struct {
	ID        string
	Amount    float64
	Status    Status
	CreatedAt time.Time
}
