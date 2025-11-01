package dto

import "github.com/google/uuid"

type PublisherDto struct {
	ID      *uuid.UUID
	Name    string `binding:"required_without=ID" json:"name"`
	Address string `binding:"required_without=ID" json:"address"`
	// TODO: validate phone start with 08
	Phone string `binding:"required_without=ID" json:"phone"`
	// TODO: validate website is valid url
	Website string `binding:"required_without=ID" json:"website"`
}
