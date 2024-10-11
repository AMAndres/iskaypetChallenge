package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func InitDB() {

	var err error

	db, err = sql.Open("postgres", getDBConfig())
	if err != nil {
		logrus.Error("Error al conectar a la base de datos:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		logrus.Error("Error al verificar la conexión:", err)
		return
	}

	logrus.Info("Conexión a la base de datos establecida")
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
