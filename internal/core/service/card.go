package service

import (
	"context"
	"paymentweb/internal/core/domain"
	"paymentweb/internal/core/ports/out"

	"github.com/google/uuid"
)

type CardService struct {
	repo out.CardRepository
}

func NewCardService(repo out.CardRepository) *CardService {
	return &CardService{repo}
}

func (cs CardService) Save(ctx *context.Context, idUser uuid.UUID, card *domain.Card) error {
	return cs.repo.Save(ctx, idUser, card)
}

func (cs CardService) List(ctx *context.Context) ([]domain.Card, error) {
	return cs.repo.List(ctx)
}
