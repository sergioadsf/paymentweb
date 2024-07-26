package request

import (
	"paymentweb/internal/core/domain"
	"strconv"
	"strings"
)

type CreateUserRequest struct {
	Name              string             `json:"name"`
	CPF               string             `json:"cpf"`
	Email             string             `json:"email" binding:"required"`
	Phone             string             `json:"phone_number" binding:"required"`
	CreateCardRequest *CreateCardRequest `json:"credit_card" binding:"required"`
}

func (c CreateUserRequest) NewUserDomain() *domain.User {
	cardRequest := c.CreateCardRequest
	// var exp []string
	exp := strings.Split(cardRequest.Expiration, "/")
	month, _ := strconv.Atoi(exp[0])
	year, _ := strconv.Atoi(exp[1])
	return &domain.User{
		Name:  c.Name,
		Cpf:   c.CPF,
		Email: c.Email,
		Phone: c.Phone,
		Card: &domain.Card{
			Brand:    cardRequest.Type,
			Number:   cardRequest.Number,
			Alias:    cardRequest.Alias,
			ExpYear:  year,
			ExpMonth: month,
		},
	}
}
