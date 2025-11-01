package dto

import "github.com/google/uuid"

type PublisherDto struct {
	ID      *uuid.UUID
	Name    string `binding:"required_without=ID" json:"name"`
	Address string `binding:"required_without=ID" json:"address"`
	Phone   string `binding:"required_without=ID,startswith=08" json:"phone"`
	Website string `binding:"required_without=ID,url" json:"website"`
}
