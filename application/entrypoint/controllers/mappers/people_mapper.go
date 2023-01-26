package mappers

import (
	"rest-api/application/domain"
	"rest-api/application/entrypoint/controllers/dto"
)

func PeopleToPeopleDto(people *domain.Person) *dto.PersonDTO {
	return &dto.PersonDTO{
		ID:        people.ID,
		Firstname: people.Firstname,
		Lastname:  people.Lastname,
	}
}

func PeopleDTOToPeople(peopleDto *dto.PersonDTO) *domain.Person {
	return &domain.Person{
		ID:        peopleDto.ID,
		Firstname: peopleDto.Firstname,
		Lastname:  peopleDto.Lastname,
	}
}
