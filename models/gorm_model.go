package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primaryKey" json:"id" form:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" form:"updated_at"`
}
