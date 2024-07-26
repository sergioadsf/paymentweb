package response

import (
	"paymentweb/internal/core/domain"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Cpf  *string   `json:"cpf,omitempty"`
}

func NewUserResponse(u domain.User) *UserResponse {
	return &UserResponse{
		*u.Id,
		u.Name,
		nil,
	}
}

func NewUserListResponse(listUser []domain.User) []UserResponse {
	var users []UserResponse
	for _, u := range listUser {
		users = append(users, *NewUserResponse(u))
	}
	return users
}
