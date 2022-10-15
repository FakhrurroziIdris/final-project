package models

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url"`
	UserID   uint   `gorm:"not null" json:"user_id" form:"user_id"`
}

// func (payload *Photo) BeforeCreate(tx *gorm.DB) error {
// 	_, err := mail.ParseAddress(payload.Email)
// 	if err != nil {
// 		return errors.New("Format email salah")
// 	}

// 	if payload.Age <= 8 {
// 		return errors.New("Age harus lebih dari 8 tahun")
// 	}

// 	if len(payload.Password) < 6 {
// 		return errors.New("Password minimal harus 6 karakter")
// 	}

// 	payload.Password = bcrypt.HashPass(payload.Password)

// 	return nil
// }
