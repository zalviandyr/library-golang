package dto

import "github.com/google/uuid"

type AuthorDto struct {
	ID        *uuid.UUID
	FirstName string `binding:"required_without=ID" json:"first_name"`
	LastName  string `binding:"required_without=ID" json:"last_name"`
	Biography string `binding:"required_without=ID" json:"biography"`
}
