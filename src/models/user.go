package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	GoogleId string
	Token    string
}
