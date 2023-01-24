package database

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var (
	// DB variable for connection DB postgresql
	DB *sql.DB
)

func DbConnect() error {
	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	// defer result.Close()
	err = result.Ping()
	if err != nil {
		log.Println("Error DB Ping : ", err)
		return err
	}

	DB = result
	return nil
}
