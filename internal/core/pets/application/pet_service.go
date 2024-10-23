package application

import (
	"context"
	"math"
	"time"

	petDomain "github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
	apiModels "github.com/AMAndres/iskaypetChallenge/models"
	"gonum.org/v1/gonum/stat"
)

var PetServiceInstance PetServiceImpl

type PetServiceImpl struct {
	petRepository *petDomain.PetRepository
}

func init() {
	PetServiceInstance = PetServiceImpl{}
}

func (o *PetServiceImpl) NewPetService(ps *petDomain.PetRepository) *PetService {

	// Idempotent behavior
	if PetServiceInstance.petRepository == nil {
		PetServiceInstance.petRepository = ps
	}

	iFace := PetService(&PetServiceInstance)

	return &iFace
}

func (o *PetServiceImpl) AddPet(ctx context.Context, pet *petDomain.Pet) (*petDomain.Pet, error) {
	return (*o.petRepository).AddPet(ctx, pet)
}

func (o *PetServiceImpl) GetAllPets(ctx context.Context) (*[]petDomain.Pet, error) {
	return (*o.petRepository).GetAllPets(ctx)
}

func (o *PetServiceImpl) GetPetAppById(ctx context.Context, pet petDomain.Pet) (*petDomain.Pet, error) {
	return (*o.petRepository).GetPetById(ctx, pet.ID)
}

func (o *PetServiceImpl) MostNumerousSpecies(ctx context.Context) (*apiModels.KpiMostNumerousSpecies, error) {
	return (*o.petRepository).MostNumerousSpecies(ctx)
}

func (o *PetServiceImpl) AverageAge(ctx context.Context, speciesName *string) (*apiModels.KpiAverageAge, error) {

	var response = apiModels.KpiAverageAge{}

	// Get average age
	avAge, ageCollection, err := averageAgeBySpecies(ctx, *speciesName, o.petRepository)
	if err != nil {
		return nil, err
	}

	// Calculate the standard deviation of age
	_, ageSD := stdDeviationAge(ageCollection)

	response.AverageAge = avAge
	response.AgeStandardDeviation = ageSD

	return &response, nil
}

func averageAgeBySpecies(ctx context.Context, speciesName string, repo *petDomain.PetRepository) (float64, *[]float64, error) {

	var err error
	var pets *[]petDomain.Pet
	var ageAggregation float64
	var ageCollection []float64

	pets, err = (*repo).GetAllPetsBySpecies(ctx, speciesName)
	if err != nil {
		return 0.0, nil, err
	}

	for _, pet := range *pets {
		diff := time.Since(pet.Birthday)
		years := float64(diff.Hours() / 24 / 365)
		ageCollection = append(ageCollection, years)
		ageAggregation += years
	}

	return (ageAggregation / float64(len(*pets))), &ageCollection, nil
}

func stdDeviationAge(ageCollection *[]float64) (float64, float64) {
	mean := stat.Mean(*ageCollection, nil)
	variance := stat.Variance(*ageCollection, nil)
	stddev := math.Sqrt(variance)

	return mean, stddev
}
