package services

import (
	"final-project/repositories"
)

type SocialMedia struct {
	repo repositories.IRepository
}

func SocialMediaService(repo repositories.IRepository) *SocialMedia {
	return &SocialMedia{repo: repo}
}

func (service *SocialMedia) UpdateSocialMedia() {
	// service.repo.UpdateSocialMedia()
}

func (service *SocialMedia) Get(payload interface{}) (interface{}, error) {
	data, err := service.repo.Get(payload)
	return data, err
}

func (service *SocialMedia) Create(payload interface{}) (interface{}, error) {
	socialMedia, err := service.repo.Create(payload)
	return socialMedia, err
}

func (service *SocialMedia) GetOne(payload interface{}) (interface{}, error) {
	data, err := service.repo.GetOne(payload)
	return data, err
}

func (service *SocialMedia) Update(payload interface{}) (interface{}, error) {
	data, err := service.repo.Update(payload)
	return data, err
}

func (service *SocialMedia) Delete(payload interface{}) (interface{}, error) {
	data, err := service.repo.Delete(payload)
	return data, err
}
