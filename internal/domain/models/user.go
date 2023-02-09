package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
	PassHash string
	Email    string
}
