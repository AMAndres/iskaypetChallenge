package config

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() error {

	var err error

	db, err = sql.Open("postgres", getDBConfig())
	if err != nil {
		return errors.New(fmt.Sprintf("Error al conectar a la base de datos. Error: %s", err.Error()))
	}

	err = db.Ping()
	if err != nil {
		return errors.New(fmt.Sprintf("Error al verificar la conexión. Error: %s", err.Error()))
	}

	return nil
}

func getDBConfig() string {

	//connStr := "postgres://iskaypet:iskaypet@127.0.0.1:5432/postgres?sslmode=disable"

	host := "localhost"
	port := 5432
	user := "iskaypet"
	password := "iskaypet"
	dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return psqlInfo
}

func GetDB() *sql.DB {
	return db
}
