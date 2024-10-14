package infraestructure

import (
	"context"

	petService "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	petDomain "github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
	petMapper "github.com/AMAndres/iskaypetChallenge/internal/core/pets/infraestructure"
	domainErrors "github.com/AMAndres/iskaypetChallenge/internal/core/transversal/domain/errors"
	operationApi "github.com/AMAndres/iskaypetChallenge/restapi/operations/pets"
	"github.com/go-openapi/runtime/middleware"
)

func GetPetsByIDHandler(params operationApi.GetPetsByIDParams) middleware.Responder {

	pet, err := petService.GetService().GetPetAppById(context.Background(), petDomain.Pet{ID: params.PetID})
	if err != nil {
		switch err.(type) {
		case *domainErrors.NotFoundError:
			return operationApi.NewGetPetsByIDNoContent()
		default:
			return operationApi.NewGetPetsByIDInternalServerError()
		}
	}

	response := petMapper.MapPetDomainToApiModel(pet)
	return operationApi.NewGetPetsByIDOK().WithPayload(response)
}
