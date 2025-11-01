package dao

import (
	"time"

	"github.com/google/uuid"
)

type BaseDao struct {
	ID        *uuid.UUID `gorm:"type:UUID;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
