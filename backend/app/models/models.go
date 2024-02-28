package models

import (
	"time"

	"github.com/google/uuid"
)

type Anime struct {
	ID          uuid.UUID  `gorm:"primaryKey" json:"id"`
	Name        *string    `json:"name"`
	ReleaseDate *time.Time `json:"release_date"`
	Gender      *string    `json:"gender"`
}
