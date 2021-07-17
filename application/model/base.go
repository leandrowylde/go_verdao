package model

import "time"

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
}
