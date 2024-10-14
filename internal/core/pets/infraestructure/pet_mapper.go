package infraestructure

import (
	"time"

	petDomain "github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
	apiModels "github.com/AMAndres/iskaypetChallenge/models"
	"github.com/go-openapi/strfmt"
)

func MapApiModelToPetDomain(input *apiModels.AddPet) *petDomain.Pet {

	output := petDomain.Pet{
		Name:     *input.Name,
		Species:  int8(*input.Species),
		Gender:   int8(*input.Gender),
		Birthday: time.Time(*input.Birthday),
	}

	return &output
}

func MapPetDomainToApiModel(input *petDomain.Pet) *apiModels.Pet {

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

func MapPetsDomainToApiModel(input *[]petDomain.Pet) []*apiModels.Pet {

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
