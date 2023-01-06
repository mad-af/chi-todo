package models

import (
	"time"
)

type Activities struct {
	ID        int        `json:"id"`
	Email     *string    `json:"email"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:current_timestamp()"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:current_timestamp()"`
	DeletedAt *time.Time `json:"-"`
}
