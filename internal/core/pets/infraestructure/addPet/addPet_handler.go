package infraestructure

import (
	"context"
	"time"

	petService "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	petDomain "github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
	apiModels "github.com/AMAndres/iskaypetChallenge/models"
	operationApi "github.com/AMAndres/iskaypetChallenge/restapi/operations/pets"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func AddPetHandler(params operationApi.AddPetParams) middleware.Responder {

	// Communicating with business logic using its own language helps keep business logic clean.
	newPet := mapApiModelToPetDomain(params.Body)

	pet, err := petService.GetService().AddPet(context.Background(), newPet)
	if err != nil {
		return operationApi.NewGetPetsByIDInternalServerError()
	}

	response := mapPetDomainToApiModel(pet)
	return operationApi.NewAddPetCreated().WithPayload(response)
}

// TODO Move to a mapper package
func mapApiModelToPetDomain(input *apiModels.AddPet) *petDomain.Pet {

	output := petDomain.Pet{
		Name:     *input.Name,
		Species:  int8(*input.Species),
		Gender:   int8(*input.Gender),
		Birthday: time.Time(*input.Birthday),
	}

	return &output
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
