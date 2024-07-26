package domain

import "github.com/google/uuid"

type Card struct {
	Id       *uuid.UUID
	Brand    string
	Number   string
	Alias    string
	ExpYear  int
	ExpMonth int
}
