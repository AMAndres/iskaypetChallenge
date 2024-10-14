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

	response := mapPetDomainToApiModel(pets)
	return operationApi.NewGetAllPetsOK().WithPayload(response)
}

// TODO Move to a mapper package
func mapPetDomainToApiModel(input *[]petDomain.Pet) []*apiModels.Pet {

	var output []*apiModels.Pet

	for _, value := range *input {
		pet := apiModels.Pet{}

		birthday := (strfmt.Date)(value.Birthday)
		gender := int64(value.Gender)
		id := value.ID
		name := value.Name
		species := int64(value.Species)

		pet.Birthday = &birthday
		pet.Gender = &gender
		pet.ID = &id
		pet.Name = &name
		pet.Species = &species

		output = append(output, &pet)
	}

	return output
}
