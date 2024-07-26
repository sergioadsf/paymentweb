package response

import (
	"paymentweb/internal/core/domain"

	"github.com/google/uuid"
)

type CardResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Cpf  *string   `json:"cpf,omitempty"`
}

func NewCardResponse(u domain.Card) *CardResponse {
	return &CardResponse{
		*u.Id,
		u.Brand,
		nil,
	}
}

func NewCardListResponse(listCard []domain.Card) []CardResponse {
	var cards []CardResponse
	for _, u := range listCard {
		cards = append(cards, *NewCardResponse(u))
	}
	return cards
}
