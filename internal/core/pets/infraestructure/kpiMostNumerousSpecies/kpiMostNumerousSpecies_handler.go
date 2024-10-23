package infraestructure

import (
	"context"

	petApp "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	domainErrors "github.com/AMAndres/iskaypetChallenge/internal/core/transversal/domain/errors"
	operationsKpi "github.com/AMAndres/iskaypetChallenge/restapi/operations/kpis"
	"github.com/go-openapi/runtime/middleware"
)

func KpiMostNumerousSpeciesHandler(params operationsKpi.MostNumerousSpeciesParams) middleware.Responder {

	response, err := petApp.PetServiceInstance.MostNumerousSpecies(context.Background())
	if err != nil {
		switch err.(type) {
		case *domainErrors.NotFoundError:
			return operationsKpi.NewMostNumerousSpeciesNoContent()
		default:
			return operationsKpi.NewMostNumerousSpeciesInternalServerError()
		}
	}

	return operationsKpi.NewMostNumerousSpeciesOK().WithPayload(response)
}
