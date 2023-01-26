package repositories

import "rest-api/application/domain"

type PeopleRepository interface {
	GetPeoples() ([]*domain.Person, error)
	GetPeople(id string) (*domain.Person, error)
	CreatePeople(people *domain.Person) (*domain.Person, error)
	DeletePeople(id string) error
}
