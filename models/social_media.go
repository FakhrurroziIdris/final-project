package models

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	UserID         uint   `gorm:"not null;" json:"user_id" form:"user_id"`
}
