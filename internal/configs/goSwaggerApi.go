package config

import (
	"log"

	"github.com/AMAndres/iskaypetChallenge/restapi"
	"github.com/AMAndres/iskaypetChallenge/restapi/operations"
	"github.com/go-openapi/loads"
)

func GoSwaggerApi() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewBackendAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8081
	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
