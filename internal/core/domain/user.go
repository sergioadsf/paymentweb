package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id    *uuid.UUID
	Name  string
	Cpf   string
	Email string
	Phone string
	Card  *Card
}
