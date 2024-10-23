package infraestructure

import (
	"context"

	petApp "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	petInfraMapper "github.com/AMAndres/iskaypetChallenge/internal/core/pets/infraestructure"
	domainErrors "github.com/AMAndres/iskaypetChallenge/internal/core/transversal/domain/errors"
	operationApi "github.com/AMAndres/iskaypetChallenge/restapi/operations/pets"
	"github.com/go-openapi/runtime/middleware"
)

func GetAllPetsHandler(params operationApi.GetAllPetsParams) middleware.Responder {

	pets, err := petApp.PetServiceInstance.GetAllPets(context.Background())
	if err != nil {
		switch err.(type) {
		case *domainErrors.NotFoundError:
			return operationApi.NewGetAllPetsNoContent()
		default:
			return operationApi.NewGetAllPetsInternalServerError()
		}
	}

	response := petInfraMapper.MapPetsDomainToApiModel(pets)
	return operationApi.NewGetAllPetsOK().WithPayload(response)
}
