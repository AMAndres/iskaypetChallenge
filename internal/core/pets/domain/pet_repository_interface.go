package domain

import (
	"context"

	apiModels "github.com/AMAndres/iskaypetChallenge/models"
)

type PetRepository interface {
	AddPet(ctx context.Context, pet *Pet) (*Pet, error)
	GetAllPets(ctx context.Context) (*[]Pet, error)
	GetPetById(ctx context.Context, id int64) (*Pet, error)
	MostNumerousSpecies(ctx context.Context) (*apiModels.KpiMostNumerousSpecies, error)
	GetAllPetsBySpecies(ctx context.Context, species string) (*[]Pet, error)
}
