package out

import (
	"context"
	"paymentweb/internal/core/domain"
)

type UserRepository interface {
	Save(ctx *context.Context, user *domain.User) error
	List(ctx *context.Context) ([]domain.User, error)
	ListUser(ctx *context.Context, id int32) (*domain.User, error)
}
