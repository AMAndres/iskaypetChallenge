package infraestructure

import (
	"context"

	petService "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	petMapper "github.com/AMAndres/iskaypetChallenge/internal/core/pets/infraestructure"
	domainErrors "github.com/AMAndres/iskaypetChallenge/internal/core/transversal/domain/errors"
	operationApi "github.com/AMAndres/iskaypetChallenge/restapi/operations/pets"
	"github.com/go-openapi/runtime/middleware"
)

func GetAllPetsHandler(params operationApi.GetAllPetsParams) middleware.Responder {

	pets, err := petService.GetService().GetAllPets(context.Background())
	if err != nil {
		switch err.(type) {
		case *domainErrors.NotFoundError:
			return operationApi.NewGetAllPetsNoContent()
		default:
			return operationApi.NewGetAllPetsInternalServerError()
		}
	}

	response := petMapper.MapPetsDomainToApiModel(pets)
	return operationApi.NewGetAllPetsOK().WithPayload(response)
}
