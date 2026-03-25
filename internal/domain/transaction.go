package domain

import "time"

type Transaction struct {
	ID        string
	Amount    float64
	Status    string
	CreatedAt time.Time
}
