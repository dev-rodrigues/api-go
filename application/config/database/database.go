package database

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/application/config"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	DB *sql.DB
)

func DbConnect() (*sql.DB, error) {
	config.InitViper()
	if DB != nil {
		return DB, nil
	}

	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	// defer result.Close()
	err = result.Ping()
	if err != nil {
		log.Println("Error DB Ping : ", err)
		return nil, err
	}

	DB = result

	return DB, nil
}
