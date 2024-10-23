package application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	"github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
	domainErrors "github.com/AMAndres/iskaypetChallenge/internal/core/transversal/domain/errors"
	apiModels "github.com/AMAndres/iskaypetChallenge/models"
	"github.com/AMAndres/iskaypetChallenge/test/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApplication(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Application Suite")
}

var (
	ctx               context.Context
	mockCtrl          *gomock.Controller
	mockPetRepository *mocks.MockPetRepository
	petService        PetService
)

var _ = Describe("PetService", func() {

	BeforeEach(func() {
		ctx = context.Background()
		mockCtrl = gomock.NewController(GinkgoT())
		mockPetRepository = mocks.NewMockPetRepository(mockCtrl)
		petService = NewPetService(mockPetRepository)
	})

	Describe("Service", func() {
		Context("Scenario - Create instance", func() {
			It("Success - Instance created", func() {
				service := NewPetService(mockPetRepository)
				Expect(service).ToNot(BeNil())
			})
		})
		Context("Scenario - Get instance", func() {
			It("Success - Instance retrieved", func() {
				service := GetService()
				Expect(service).ToNot(BeNil())
			})
		})
	})

	Describe("AddPet()", func() {
		Context("Scenario - Adding some regular pet", func() {

			dPet := getDomainPet()
			dbPet := getDatabasePet()

			It("Success - Pet added successfully", func() {

				mockPetRepository.EXPECT().AddPet(ctx, dPet).Return(dbPet, nil).Times(1)
				serviceResponse, serviceErr := petService.AddPet(ctx, dPet)

				Expect(serviceErr).To(BeNil())
				Expect(serviceResponse).To(Equal(dbPet))
			})

			It("Error - Repository fails when adding a pet", func() {

				dbError := errors.New("Any error")
				expectedErr := dbError

				mockPetRepository.EXPECT().AddPet(ctx, dPet).Return(nil, dbError).Times(1)
				serviceResponse, serviceErr := petService.AddPet(ctx, dPet)

				Expect(serviceResponse).To(BeNil())
				Expect(serviceErr).To(Equal(expectedErr))
			})
		})
	})

	Describe("GetAllPets()", func() {
		Context("Scenario - Database has some pets", func() {

			dbPets := getDatabasePets()

			It("Success - DB returns one or more pets", func() {

				mockPetRepository.EXPECT().GetAllPets(ctx).Return(dbPets, nil).Times(1)
				serviceResponse, serviceErr := petService.GetAllPets(ctx)

				Expect(serviceErr).To(BeNil())
				Expect(serviceResponse).To(Equal(dbPets))
			})

			It("Error - Repository fails retrieving pet", func() {

				expectedErr := errors.New("Any error")

				mockPetRepository.EXPECT().GetAllPets(ctx).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.GetAllPets(ctx)

				Expect(serviceResponse).To(BeNil())
				Expect(serviceErr).To(Equal(expectedErr))
			})
		})
		Context("Scenario - Database is emmpty", func() {
			It("Success - DB returns one or more pets", func() {

				expectedErr := domainErrors.NewNotFoundError("pets not found")

				mockPetRepository.EXPECT().GetAllPets(ctx).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.GetAllPets(ctx)

				Expect(serviceResponse).To(BeNil())
				Expect(serviceErr).To(Equal(expectedErr))
			})

			It("Error - Repository fails retrieving pet", func() {

				expectedErr := errors.New("Any error")

				mockPetRepository.EXPECT().GetAllPets(ctx).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.GetAllPets(ctx)

				Expect(serviceResponse).To(BeNil())
				Expect(serviceErr).To(Equal(expectedErr))
			})
		})
	})

	Describe("GetPetAppById()", func() {
		Context("Scenario - Retrieving a pet by ID", func() {

			id := int64(1)
			dbPet := getDatabasePet()

			It("Success - DB returns the Pet associated to the ID provided", func() {

				mockPetRepository.EXPECT().GetPetById(ctx, id).Return(dbPet, nil).Times(1)
				serviceResponse, serviceErr := petService.GetPetAppById(ctx, *dbPet)

				Expect(serviceErr).To(BeNil())
				Expect(serviceResponse).To(Equal(dbPet))
			})

			It("Sucess - DB does not have record for the ID provided", func() {

				expectedErr := domainErrors.NewNotFoundError("pet not found")

				mockPetRepository.EXPECT().GetPetById(ctx, id).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.GetPetAppById(ctx, *dbPet)

				Expect(serviceErr).To(Equal(expectedErr))
				Expect(serviceResponse).To(BeNil())
			})

			It("Error - Repository fails retrieving pet", func() {

				expectedErr := errors.New("Any error")

				mockPetRepository.EXPECT().GetPetById(ctx, id).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.GetPetAppById(ctx, *dbPet)

				Expect(serviceErr).To(Equal(expectedErr))
				Expect(serviceResponse).To(BeNil())
			})
		})
	})

	Describe("MostNumerousSpecies()", func() {
		Context("Scenario - Database has some pets", func() {
			It("Success - Returns data related to the most numerous species", func() {

				expectedSpecies := int64(0)
				expectedSpeciesDesc := "Dog"
				expectedQty := int64(3)

				dbResponse := apiModels.KpiMostNumerousSpecies{
					MostNumerousSpecies:            &expectedSpecies,
					MostNumerousSpeciesDescription: expectedSpeciesDesc,
					MostNumerousSpeciesQty:         &expectedQty,
				}

				mockPetRepository.EXPECT().MostNumerousSpecies(ctx).Return(&dbResponse, nil).Times(1)
				serviceResponse, serviceErr := petService.MostNumerousSpecies(ctx)

				Expect(serviceResponse.MostNumerousSpecies).To(Equal(dbResponse.MostNumerousSpecies))
				Expect(serviceResponse.MostNumerousSpeciesDescription).To(Equal(dbResponse.MostNumerousSpeciesDescription))
				Expect(serviceResponse.MostNumerousSpeciesQty).To(Equal(dbResponse.MostNumerousSpeciesQty))
				Expect(serviceErr).To(BeNil())
			})

			It("Error - Some error with database", func() {

				expectedErr := errors.New("Any error")

				mockPetRepository.EXPECT().MostNumerousSpecies(ctx).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.MostNumerousSpecies(ctx)

				Expect(serviceResponse).To(BeNil())
				Expect(serviceErr).To(Equal(expectedErr))
			})
		})

		Context("Scenario - Database is empty", func() {
			It("Sucess - DB does not have records", func() {

				expectedErr := domainErrors.NewNotFoundError("species not found")

				mockPetRepository.EXPECT().MostNumerousSpecies(ctx).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.MostNumerousSpecies(ctx)

				Expect(serviceResponse).To(BeNil())
				Expect(serviceErr).To(Equal(expectedErr))
			})
		})
	})

	Describe("AverageAge()", func() {
		Context("Scenario - Database has some pets", func() {
			It("Success - Returns average age for the required species", func() {

				speciesName := "Cat"
				birthday1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
				birthday2 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
				expectedAge := 3.3
				expectedAgeStdDeviation := 2.12

				mockPetRepository.EXPECT().GetAllPetsBySpecies(ctx, speciesName).Return(&[]domain.Pet{
					{Birthday: birthday1},
					{Birthday: birthday2},
				}, nil).Times(1)
				serviceResponse, serviceErr := petService.AverageAge(ctx, &speciesName)

				Expect(serviceResponse.AverageAge).Should(BeNumerically("~", expectedAge, 0.01))
				Expect(serviceResponse.AgeStandardDeviation).Should(BeNumerically("~", expectedAgeStdDeviation, 0.01))
				Expect(serviceErr).To(BeNil())
			})

			It("Error - Error accesing db", func() {

				speciesName := "Dog"
				expectedErr := errors.New("Any error")

				mockPetRepository.EXPECT().GetAllPetsBySpecies(ctx, speciesName).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.AverageAge(ctx, &speciesName)

				Expect(serviceResponse).To(BeNil())
				Expect(serviceErr).To(Equal(expectedErr))
			})
		})

		Context("Scenario - Database is empty", func() {
			It("Sucess - DB does not have records", func() {

				speciesName := "Dog"
				expectedErr := domainErrors.NewNotFoundError("species not found")

				mockPetRepository.EXPECT().GetAllPetsBySpecies(ctx, speciesName).Return(nil, expectedErr).Times(1)
				serviceResponse, serviceErr := petService.AverageAge(ctx, &speciesName)

				Expect(serviceResponse).To(BeNil())
				Expect(serviceErr).To(Equal(expectedErr))
			})
		})
	})
})

func getDomainPet() *domain.Pet {
	pet := &domain.Pet{
		ID:       1,
		Name:     "Firulais",
		Species:  1,
		Birthday: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return pet
}

func getDatabasePet() *domain.Pet {
	pet := getDomainPet()
	pet.ID = 1
	return pet
}

func getDatabasePets() *[]domain.Pet {
	response := make([]domain.Pet, 2)

	response[0] = *getDomainPet()
	response[0].ID = 1

	response[1] = *getDomainPet()
	response[1].ID = 2

	return &response
}
