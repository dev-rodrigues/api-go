package config

import (
	"database/sql"
	"fmt"
)

func DbConnect() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "Postgres2022!"
		dbname   = "postgres"
	)

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
