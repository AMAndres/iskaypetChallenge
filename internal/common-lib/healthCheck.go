package commonlib

import (
	"fmt"
	"net/http"

	config "github.com/AMAndres/iskaypetChallenge/internal/configs"
	"github.com/sirupsen/logrus"
)

func InitHealthcheck() {
	// TODO: Store port as OS env variable
	http.HandleFunc("/health", healthCheckHandler)
	http.ListenAndServe(":8082", nil)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	if err := checkDatabaseConnection(); err != nil {
		logrus.Error(fmt.Sprintf("Error /health: %v", err.Error()))
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func checkDatabaseConnection() error {

	var err error

	if err = config.GetDB().Ping(); err != nil {
		logrus.Error("DB not responding")
	}

	return err
}
