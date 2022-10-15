package repositories

import (
	"final-project/models"
	"reflect"

	"gorm.io/gorm"
)

type socialMedia struct {
	db *gorm.DB
}

func SocialMediaRepository(db *gorm.DB) *socialMedia {
	return &socialMedia{db: db}
}

func (repo *socialMedia) UpdateSensor() {
	// repo.db.Update()
}

func (repo *socialMedia) Get(payload interface{}) (interface{}, error) {
	criteria := reflect.ValueOf(payload).Interface().(models.SocialMedia)
	response := []models.SocialMedia{}
	user := models.User{}

	err := repo.db.Where("user_id = ?", criteria.UserID).Find(&response).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Preload("Photos").Where("id = ?", criteria.UserID).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"social_medias": response,
		"user":          user,
	}, nil
}

func (repo *socialMedia) Create(payload interface{}) (interface{}, error) {
	tx := repo.db.Begin()
	socialMedia := reflect.ValueOf(payload).Interface().(models.SocialMedia)

	err := tx.Create(&socialMedia).Error
	if err != nil {
		tx.Rollback()
		return payload, err
	}
	tx.Commit()
	return socialMedia, err
}

func (repo *socialMedia) GetOne(payload interface{}) (interface{}, error) {
	response := reflect.ValueOf(payload).Interface().(models.SocialMedia)

	err := repo.db.Where("user_id = ? AND id = ?", response.UserID, response.ID).Take(&response).Error
	return response, err
}

func (repo *socialMedia) Update(payload interface{}) (interface{}, error) {
	socialMedia := reflect.ValueOf(payload).Interface().(models.SocialMedia)

	err := repo.db.Model(&socialMedia).Where("user_id = ? AND id = ?", socialMedia.UserID, socialMedia.ID).Updates(socialMedia).Take(&socialMedia).Error
	return socialMedia, err
}

func (repo *socialMedia) Delete(payload interface{}) (interface{}, error) {
	socialMedia := reflect.ValueOf(payload).Interface().(models.SocialMedia)

	tx := repo.db.Begin()
	deleted := repo.db.Model(&socialMedia).Where("user_id = ? AND id = ?", socialMedia.UserID, socialMedia.ID).Delete(&socialMedia)
	if deleted.Error != nil {
		tx.Rollback()
		return deleted, deleted.Error
	}

	tx.Commit()
	return deleted, nil
}
