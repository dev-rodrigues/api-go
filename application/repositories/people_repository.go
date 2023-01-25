package repositories

import "rest-api/application/domain"

type PeopleRepository interface {
	GetPeoples() ([]*domain.Person, error)
}
