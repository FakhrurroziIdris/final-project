package repositories

import (
	"final-project/models"
	"reflect"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *user {
	return &user{db: db}
}

func (repo *user) UpdateSensor() {
	// repo.db.Update()
}

func (repo *user) Get(payload interface{}) (interface{}, error) {
	response := []models.User{}
	err := repo.db.Find(&response).Error
	return response, err
}

func (repo *user) Create(payload interface{}) (interface{}, error) {
	tx := repo.db.Begin()
	user := reflect.ValueOf(payload).Interface().(models.User)

	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return payload, err
	}
	tx.Commit()
	return user, err
}

func (repo *user) GetOne(payload interface{}) (interface{}, error) {
	response := reflect.ValueOf(payload).Interface().(models.User)

	err := repo.db.Where("email = ?", response.Email).Take(&response).Error
	return response, err
}

func (repo *user) Update(payload interface{}) (interface{}, error) {
	user := reflect.ValueOf(payload).Interface().(models.User)

	err := repo.db.Model(&user).Where("email=?", user.Email).Updates(user).Take(&user).Error
	return user, err
}

func (repo *user) Delete(payload interface{}) (interface{}, error) {
	user := reflect.ValueOf(payload).Interface().(models.User)
	comment := models.Comment{}
	photo := models.Photo{}

	tx := repo.db.Begin()
	errComment := repo.db.Model(&comment).Where("user_id = ?", user.ID).Delete(&comment).Error
	if errComment != nil {
		tx.Rollback()
		return payload, errComment
	}

	errPhoto := repo.db.Model(&photo).Where("user_id=?", user.ID).Delete(photo).Error
	if errPhoto != nil {
		tx.Rollback()
		return payload, errPhoto
	}

	errUser := repo.db.Model(&user).Where("email=?", user.Email).Delete(user).Error
	if errUser != nil {
		tx.Rollback()
		return payload, errUser
	}

	tx.Commit()
	return user, nil
}
