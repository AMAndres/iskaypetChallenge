package application

import (
	"context"

	petDomain "github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
	apiModels "github.com/AMAndres/iskaypetChallenge/models"
)

type PetService interface {
	AddPet(ctx context.Context, pet *petDomain.Pet) (*petDomain.Pet, error)
	GetAllPets(ctx context.Context) (*[]petDomain.Pet, error)
	GetPetAppById(ctx context.Context, pet petDomain.Pet) (*petDomain.Pet, error)
	MostNumerousSpecies(ctx context.Context) (*apiModels.KpiMostNumerousSpecies, error)
	AverageAge(ctx context.Context, speciesName *string) (*apiModels.KpiAverageAge, error)
}
