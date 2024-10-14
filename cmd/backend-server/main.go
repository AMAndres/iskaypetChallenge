package main

import (
	"log"

	commonlib "github.com/AMAndres/iskaypetChallenge/internal/common-lib"
	config "github.com/AMAndres/iskaypetChallenge/internal/configs"
	"github.com/sirupsen/logrus"
)

func main() {

	log.Println(" -------------- Starting server --------------")

	// Load
	config.LoadConfig()
	log.Println(" ** Config loaded")

	// Set up logs
	commonlib.InitLog()
	logrus.Info(" ** Logs initiated")

	// Database
	err := config.InitDatabase()
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.Info(" ** Database loaded")

	// Healthcheck
	go commonlib.InitHealthcheck()
	logrus.Info(" ** Healthcheck loaded")

	// Dependencies injection
	config.InjectDependencies()
	logrus.Info(" ** Dependencies loaded")

	// REST API
	go config.StartGoSwaggerApi()
	logrus.Info(" ** API rest service loaded")

	// Keep server running
	commonlib.KeepServerAlive()
}
