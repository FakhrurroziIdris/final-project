package models

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoUrl string    `gorm:"not null" json:"photo_url" form:"photo_url"`
	UserID   uint      `gorm:"not null;" json:"user_id" form:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
