package config

import (
	petApp "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	petInfra "github.com/AMAndres/iskaypetChallenge/internal/core/pets/infraestructure"
)

func InjectDependencies() {

	// Repositories
	petRepo := petInfra.PetRepositoryPgInstance.NewRepository(GetDB())

	// Services
	_ = petApp.PetServiceInstance.NewPetService(petRepo)
}
