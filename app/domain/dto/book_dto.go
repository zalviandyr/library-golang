package dto

import (
	"time"

	"github.com/google/uuid"
)

type BookDto struct {
	ID          *uuid.UUID
	Title       string      `binding:"required_without=ID" json:"title"`
	Description string      `binding:"required_without=ID" json:"description"`
	ISBN        string      `binding:"required_without=ID" json:"isbn"`
	PublishedAt time.Time   `binding:"required_without=ID" json:"published_at"`
	PublisherID *uuid.UUID  `binding:"required_without=ID" json:"publisher_id"`
	AuthorIDs   []uuid.UUID `binding:"required_without=ID" json:"author_ids"`
}
