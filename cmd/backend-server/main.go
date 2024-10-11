package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AMAndres/iskaypetChallenge/internal/commonlib"
	config "github.com/AMAndres/iskaypetChallenge/internal/configs"
	"github.com/AMAndres/iskaypetChallenge/restapi"
	"github.com/AMAndres/iskaypetChallenge/restapi/operations"
	"github.com/go-openapi/loads"
)

func main() {

	log.Println(" -------------- Starting server --------------")

	// Load
	config.LoadConfig()

	// Set up logs
	commonlib.InitLog()

	// Database
	config.InitDB()

	// Healthcheck
	commonlib.InitHealthcheck()

	// Dependencies injection
	config.InjectDependencies()

	// REST API
	go goSwaggerApi()

	// Keep server running
	keepAliveServer()
}

func goSwaggerApi() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewBackendAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func keepAliveServer() {

	// Waits for SIGTERM signal on a channel
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
