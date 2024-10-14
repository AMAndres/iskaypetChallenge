package infraestructure

import (
	"context"

	petService "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	petDomain "github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
	domainErrors "github.com/AMAndres/iskaypetChallenge/internal/core/transversal/domain/errors"
	apiModels "github.com/AMAndres/iskaypetChallenge/models"
	operationApi "github.com/AMAndres/iskaypetChallenge/restapi/operations/pets"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
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

	response := mapPetDomainToApiModel(pet)
	return operationApi.NewGetPetsByIDOK().WithPayload(response)
}

// TODO Move to a mapper package
func mapPetDomainToApiModel(input *petDomain.Pet) *apiModels.Pet {

	birthday := (strfmt.Date)(input.Birthday)
	gender := int64(input.Gender)
	id := input.ID
	name := input.Name
	species := int64(input.Species)

	output := apiModels.Pet{
		Birthday: &birthday,
		Gender:   &gender,
		ID:       &id,
		Name:     &name,
		Species:  &species,
	}

	return &output
}
