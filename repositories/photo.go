package repositories

import (
	"final-project/models"
	"reflect"

	"gorm.io/gorm"
)

type photo struct {
	db *gorm.DB
}

func PhotoRepository(db *gorm.DB) *photo {
	return &photo{db: db}
}

func (repo *photo) UpdateSensor() {
	// repo.db.Update()
}

func (repo *photo) Get(payload interface{}) (interface{}, error) {
	criteria := reflect.ValueOf(payload).Interface().(models.Photo)
	response := []models.Photo{}

	err := repo.db.Where("user_id = ?", criteria.UserID).Find(&response).Error
	return response, err
}

func (repo *photo) Create(payload interface{}) (interface{}, error) {
	tx := repo.db.Begin()
	photo := reflect.ValueOf(payload).Interface().(models.Photo)

	err := tx.Create(&photo).Error
	if err != nil {
		tx.Rollback()
		return payload, err
	}
	tx.Commit()
	return photo, err
}

func (repo *photo) GetOne(payload interface{}) (interface{}, error) {
	response := reflect.ValueOf(payload).Interface().(models.Photo)

	err := repo.db.Where("user_id = ? AND id = ?", response.UserID, response.ID).Take(&response).Error
	return response, err
}

func (repo *photo) Update(payload interface{}) (interface{}, error) {
	photo := reflect.ValueOf(payload).Interface().(models.Photo)

	err := repo.db.Model(&photo).Where("user_id = ? AND id = ?", photo.UserID, photo.ID).Updates(photo).Take(&photo).Error
	return photo, err
}

func (repo *photo) Delete(payload interface{}) (interface{}, error) {
	photo := reflect.ValueOf(payload).Interface().(models.Photo)
	comment := models.Comment{}

	tx := repo.db.Begin()
	commentDeleted := repo.db.Model(&comment).Where("user_id = ? AND photo_id = ?", photo.UserID, photo.ID).Delete(&comment)
	if commentDeleted.Error != nil {
		tx.Rollback()
		return commentDeleted, commentDeleted.Error
	}

	deleted := repo.db.Model(&photo).Where("user_id = ? AND id = ?", photo.UserID, photo.ID).Delete(&photo)
	if deleted.Error != nil {
		tx.Rollback()
		return deleted, deleted.Error
	}

	tx.Commit()
	return deleted, nil
}
