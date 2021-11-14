package models

import "time"

type Model struct {
	ID        uint64    `gorm:"column:id" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
}
