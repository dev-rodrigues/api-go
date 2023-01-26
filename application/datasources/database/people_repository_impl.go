package database

import (
	"database/sql"
	"rest-api/application/domain"
)

func NewPeopleRepository(db *sql.DB) *DB {
	return &DB{Postgres: db}
}

type DB struct {
	Postgres *sql.DB
}

func (db *DB) GetPeoples() ([]*domain.Person, error) {
	rows, err := (*db.Postgres).Query("SELECT * FROM people")

	if err != nil {
		return nil, err
	}

	var peoples = make([]*domain.Person, 0)

	for rows.Next() {
		var person domain.Person
		err = rows.Scan(&person.ID, &person.Firstname, &person.Lastname)
		if err != nil {
			return nil, err
		}
		peoples = append(peoples, &person)
	}
	return peoples, nil
}

func (db *DB) GetPeople(id string) (*domain.Person, error) {
	var person domain.Person
	err := (*db.Postgres).QueryRow("SELECT * FROM people WHERE id = $1", id).Scan(&person.ID, &person.Firstname, &person.Lastname)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (db *DB) CreatePeople(people *domain.Person) (*domain.Person, error) {
	err := (*db.Postgres).QueryRow("INSERT INTO people (id, firstname, lastname) VALUES ($1, $2, $3) RETURNING id", people.ID, people.Firstname, people.Lastname).Scan(&people.ID)
	if err != nil {
		return nil, err
	}
	return people, nil
}

func (db *DB) DeletePeople(id string) error {
	response, error := (*db.Postgres).Query("DELETE FROM people WHERE id = $1", id)

	if error != nil {
		return error
	}

	defer response.Close()
	return nil
}
