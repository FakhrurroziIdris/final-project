package services

import (
	"final-project/repositories"
)

type Comment struct {
	repo repositories.IRepository
}

func CommentService(repo repositories.IRepository) *Comment {
	return &Comment{repo: repo}
}

func (service *Comment) UpdateComment() {
	// service.repo.UpdateComment()
}

func (service *Comment) Get(payload interface{}) (interface{}, error) {
	data, err := service.repo.Get(payload)
	return data, err
}

func (service *Comment) Create(payload interface{}) (interface{}, error) {
	comment, err := service.repo.Create(payload)
	return comment, err
}

func (service *Comment) GetOne(payload interface{}) (interface{}, error) {
	data, err := service.repo.GetOne(payload)
	return data, err
}

func (service *Comment) Update(payload interface{}) (interface{}, error) {
	data, err := service.repo.Update(payload)
	return data, err
}

func (service *Comment) Delete(payload interface{}) (interface{}, error) {
	data, err := service.repo.Delete(payload)
	return data, err
}
