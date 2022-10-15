package repositories

import (
	"final-project/models"
	"reflect"

	"gorm.io/gorm"
)

type comment struct {
	db *gorm.DB
}

func CommentRepository(db *gorm.DB) *comment {
	return &comment{db: db}
}

func (repo *comment) UpdateSensor() {
	// repo.db.Update()
}

func (repo *comment) Get(payload interface{}) (interface{}, error) {
	criteria := reflect.ValueOf(payload).Interface().(models.Comment)
	comments := []models.Comment{}
	photo := models.Photo{}
	user := models.User{}

	err := repo.db.Where("user_id = ?", criteria.UserID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	errPhoto := repo.db.Where("user_id = ? and id = ?", comments[0].UserID, comments[0].PhotoID).First(&photo).Error
	if errPhoto != nil {
		return nil, errPhoto
	}

	errUser := repo.db.Where("id = ?", photo.UserID).First(&user).Error
	if errUser != nil {
		return nil, errUser
	}

	response := map[string]interface{}{
		"comments": comments,
		"photo":    photo,
		"user":     user,
	}
	return response, nil
}

func (repo *comment) Create(payload interface{}) (interface{}, error) {
	tx := repo.db.Begin()
	comment := reflect.ValueOf(payload).Interface().(models.Comment)

	err := tx.Create(&comment).Error
	if err != nil {
		tx.Rollback()
		return payload, err
	}
	tx.Commit()
	return comment, err
}

func (repo *comment) GetOne(payload interface{}) (interface{}, error) {
	response := reflect.ValueOf(payload).Interface().(models.Comment)

	err := repo.db.Where("user_id = ? AND id = ?", response.UserID, response.ID).Take(&response).Error
	return response, err
}

func (repo *comment) Update(payload interface{}) (interface{}, error) {
	comment := reflect.ValueOf(payload).Interface().(models.Comment)
	photo := models.Photo{}

	err := repo.db.Model(&comment).Where("user_id = ? AND id = ?", comment.UserID, comment.ID).Updates(comment).Take(&comment).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Model(&photo).Where("user_id = ? AND id = ?", comment.UserID, comment.PhotoID).Take(&photo).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"comment": comment,
		"photo":   photo,
	}, err
}

func (repo *comment) Delete(payload interface{}) (interface{}, error) {
	comment := reflect.ValueOf(payload).Interface().(models.Comment)

	tx := repo.db.Begin()
	commentDeleted := repo.db.Model(&comment).Where("user_id = ? AND id = ?", comment.UserID, comment.ID).Delete(&comment)
	if commentDeleted.Error != nil {
		tx.Rollback()
		return commentDeleted, commentDeleted.Error
	}

	tx.Commit()
	return commentDeleted, nil
}
