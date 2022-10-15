package services

import (
	"final-project/repositories"
)

type User struct {
	repo repositories.IRepository
}

func UserService(repo repositories.IRepository) *User {
	return &User{repo: repo}
}

func (service *User) UpdateUser() {
	// service.repo.UpdateUser()
}

func (service *User) Get() (interface{}, error) {
	data, err := service.repo.Get()
	return data, err
}

func (service *User) Create(payload interface{}) (interface{}, error) {
	user, err := service.repo.Create(payload)
	return user, err
}

func (service *User) GetOne(payload interface{}) (interface{}, error) {
	data, err := service.repo.GetOne(payload)
	return data, err
}

func (service *User) Update(payload interface{}) (interface{}, error) {
	data, err := service.repo.Update(payload)
	return data, err
}

func (service *User) Delete(payload interface{}) (interface{}, error) {
	data, err := service.repo.Delete(payload)
	return data, err
}
