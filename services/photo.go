package services

import (
	"final-project/repositories"
)

type Photo struct {
	repo repositories.IRepository
}

func PhotoService(repo repositories.IRepository) *Photo {
	return &Photo{repo: repo}
}

func (service *Photo) UpdatePhoto() {
	// service.repo.UpdatePhoto()
}

func (service *Photo) Get(payload interface{}) (interface{}, error) {
	data, err := service.repo.Get(payload)
	return data, err
}

func (service *Photo) Create(payload interface{}) (interface{}, error) {
	photo, err := service.repo.Create(payload)
	return photo, err
}

func (service *Photo) GetOne(payload interface{}) (interface{}, error) {
	data, err := service.repo.GetOne(payload)
	return data, err
}

func (service *Photo) Update(payload interface{}) (interface{}, error) {
	data, err := service.repo.Update(payload)
	return data, err
}

func (service *Photo) Delete(payload interface{}) (interface{}, error) {
	data, err := service.repo.Delete(payload)
	return data, err
}
