package models

import (
	"time"
)

type Todos struct {
	ID              int        `json:"id"`
	ActivityGroupID int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        bool       `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       time.Time  `json:"created_at" gorm:"default:current_timestamp()"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"default:current_timestamp()"`
	DeletedAt       *time.Time `json:"-"`
}
