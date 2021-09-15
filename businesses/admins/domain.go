package admins

import (
	"context"
	"peduli-covid/businesses/users"
)

type Usecase interface {
	Login(ctx context.Context, email, password string) (string, error)
	Store(ctx context.Context, data *users.Domain) error
}
