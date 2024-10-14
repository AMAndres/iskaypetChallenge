package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() error {

	var err error

	db, err = sql.Open("postgres", getDBConfig())
	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos. Error: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error al verificar la conexión. Error: %s", err.Error())
	}

	return nil
}

func getDBConfig() string {

	// TODO: Store properties as secrets via cloud provider / argocd / kubernetes / etc
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "admin"
	dbname := "iskaypetchallenge"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return psqlInfo
}

func GetDB() *sql.DB {
	return db
}
