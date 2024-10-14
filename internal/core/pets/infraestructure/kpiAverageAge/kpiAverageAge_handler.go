package infraestructure

import (
	"context"

	kpiService "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	operationsKpi "github.com/AMAndres/iskaypetChallenge/restapi/operations/kpis"
	"github.com/go-openapi/runtime/middleware"
)

func KpiAverageAgeHandler(params operationsKpi.AverageAgeParams) middleware.Responder {

	response, err := kpiService.GetService().AverageAge(context.Background(), params.SpeciesName)
	if err != nil {
		return operationsKpi.NewMostNumerousSpeciesInternalServerError()
	}

	return operationsKpi.NewAverageAgeOK().WithPayload(response)
}
