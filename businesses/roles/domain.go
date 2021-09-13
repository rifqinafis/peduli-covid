package roles

import (
	"time"
)

type Domain struct {
	ID        int       `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Usecase interface {
	GetByID(id int) string
}

type Repository interface {
	GetByID(id int) (Domain, error)
}
