package impl

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
