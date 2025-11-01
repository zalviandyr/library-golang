package dao

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	BaseDao

	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	ISBN        string    `gorm:"not null" json:"isbn"`
	PublishedAt time.Time `gorm:"not null" json:"published_at"`

	PublisherID *uuid.UUID `gorm:"not null;index" json:"publisher_id"`
	Publisher   *Publisher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"publisher"`

	Author []Author `gorm:"many2many:author_books;constraint:OnDelete:CASCADE,OnDelete:CASCADE" json:"authors"`
}
