package models

import (
	"errors"
	"net/mail"

	"gorm.io/gorm"

	bcrypt "final-project/utils"
)

type User struct {
	GormModel
	UserName string `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Age      int    `gorm:"not null" json:"age" form:"age"`
}

func (payload *User) BeforeCreate(tx *gorm.DB) error {
	// _, errCreate := govalidator.ValidateStruct(u)

	// if errCreate != nil {
	// 	err = errCreate
	// 	return
	// }

	_, err := mail.ParseAddress(payload.Email)
	if err != nil {
		return errors.New("Format email salah")
	}

	if payload.Age <= 8 {
		return errors.New("Age harus lebih dari 8 tahun")
	}

	if len(payload.Password) < 6 {
		return errors.New("Password minimal harus 6 karakter")
	}

	payload.Password = bcrypt.HashPass(payload.Password)

	return nil
}
