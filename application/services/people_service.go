package services

import (
	"rest-api/application/domain"
	"rest-api/application/repositories"
)

func NewPeopleService(repository repositories.PeopleRepository) *PeopleServiceDependencies {
	return &PeopleServiceDependencies{repository: repository}
}

type PeopleServiceDependencies struct {
	repository repositories.PeopleRepository
}

func (s PeopleServiceDependencies) GetPeoples() ([]*domain.Person, error) {
	return s.repository.GetPeoples()
}
