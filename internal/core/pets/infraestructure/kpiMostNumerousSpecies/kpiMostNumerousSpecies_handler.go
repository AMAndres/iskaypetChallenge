package infraestructure

import (
	"context"

	kpiService "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	operationsKpi "github.com/AMAndres/iskaypetChallenge/restapi/operations/kpis"
	"github.com/go-openapi/runtime/middleware"
)

func KpiMostNumerousSpeciesHandler(params operationsKpi.MostNumerousSpeciesParams) middleware.Responder {

	response, err := kpiService.GetService().MostNumerousSpecies(context.Background())
	if err != nil {
		return operationsKpi.NewMostNumerousSpeciesInternalServerError()
	}

	return operationsKpi.NewMostNumerousSpeciesOK().WithPayload(response)
}
