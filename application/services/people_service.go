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

func (s PeopleServiceDependencies) GetPeople(id string) (*domain.Person, error) {
	return s.repository.GetPeople(id)
}

func (s PeopleServiceDependencies) CreatePeople(people *domain.Person) (*domain.Person, error) {
	return s.repository.CreatePeople(people)
}

func (s PeopleServiceDependencies) DeletePeople(id string) error {
	return s.repository.DeletePeople(id)
}
