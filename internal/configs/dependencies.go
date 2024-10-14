package config

import (
	petService "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	petsInfraestructure "github.com/AMAndres/iskaypetChallenge/internal/core/pets/infraestructure"
)

func InjectDependencies() {

	// GetPets Service
	petService.NewPetService(petsInfraestructure.NewPostgresPetRepository(GetDB()))
}
