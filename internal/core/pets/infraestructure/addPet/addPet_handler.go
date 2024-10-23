package infraestructure

import (
	"context"

	petApp "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	petMapper "github.com/AMAndres/iskaypetChallenge/internal/core/pets/infraestructure"
	operationApi "github.com/AMAndres/iskaypetChallenge/restapi/operations/pets"
	"github.com/go-openapi/runtime/middleware"
)

func AddPetHandler(params operationApi.AddPetParams) middleware.Responder {

	// Communicating with business logic using its own language helps keep business logic clean.
	newPet := petMapper.MapApiModelToPetDomain(params.Body)

	pet, err := petApp.PetServiceInstance.AddPet(context.Background(), newPet)
	if err != nil {
		return operationApi.NewGetPetsByIDInternalServerError()
	}

	response := petMapper.MapPetDomainToApiModel(pet)
	return operationApi.NewAddPetCreated().WithPayload(response)
}
