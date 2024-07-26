package in

import (
	"context"
	"paymentweb/internal/core/domain"

	"github.com/google/uuid"
)

type CardService interface {
	Save(ctx *context.Context, idUser uuid.UUID, card *domain.Card) error
	List(ctx *context.Context) ([]domain.Card, error)
}
