package roles

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int            `json:"id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	GetByID(id int) string
}

type Repository interface {
	GetByID(id int) (Domain, error)
}
