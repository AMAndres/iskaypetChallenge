package application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/AMAndres/iskaypetChallenge/internal/core/pets/application"
	"github.com/AMAndres/iskaypetChallenge/internal/core/pets/domain"
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
	petService        *PetServiceImpl
)

var _ = Describe("PetService", func() {

	BeforeEach(func() {
		ctx = context.Background()
		mockCtrl = gomock.NewController(GinkgoT())
		mockPetRepository = mocks.NewMockPetRepository(mockCtrl)
		petService = NewPetService(mockPetRepository)
	})

	// AddPet()
	Context("Scenario - Adding some regular pet", func() {

		dPet := getDomainPet()
		dbPet := getDatabasePet()

		It("Success", func() {

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

	/*
		Describe("GetPetAppById", func() {
			Context("Scenario - Database has some pets", func() {
				It("Success - The pet associated to the ID provided is retrieved", func() {

					expectedPetID := int64(1)
					birthday, _ := time.Parse("2022-01-01", "2006-01-02")
					expectedResponse := &domain.Pet{
						ID:       expectedPetID,
						Name:     "Firulais",
						Species:  0,
						Birthday: birthday,
					}

					mockPetRepository.On("GetPetById", ctx, expectedPetID).Return(expectedResponse, nil).Once()

					serviceResponse, err := petService.GetPetAppById(ctx, domain.Pet{ID: expectedPetID})
					Expect(err).To(BeNil())
					Expect(serviceResponse).To(Equal(expectedResponse))
				})

				It("Error - Repository fails retrieving pet", func() {

					expectedPetID := int64(1)
					expectedErr := errors.New("Any error")

					mockPetRepository.On("GetPetById", ctx, expectedPetID).Return(nil, expectedErr).Once()

					_, err := petService.GetPetAppById(ctx, domain.Pet{ID: expectedPetID})
					Expect(err).To(Equal(expectedErr))
				})
			})
		})

		Describe("MostNumerousSpecies", func() {
			Context("Scenario - Database has some pets", func() {
				It("Success - Returns data related to the most numerous species", func() {

					mns := int64(0)
					mnsqty := int64(3)

					expectedSpecies := 3
					expectedSpeciesDesc := "Dog"
					expectedQty := 3

					mockPetRepository.On("MostNumerousSpecies", ctx).Return(&apiModels.KpiMostNumerousSpecies{
						MostNumerousSpecies:            &mns,
						MostNumerousSpeciesDescription: "Dog",
						MostNumerousSpeciesQty:         &mnsqty,
					}, nil).Once()

					serviceResponse, err := petService.MostNumerousSpecies(ctx)
					Expect(err).To(BeNil())
					Expect(serviceResponse.MostNumerousSpecies).To(Equal(expectedSpecies))
					Expect(serviceResponse.MostNumerousSpeciesDescription).To(Equal(expectedSpeciesDesc))
					Expect(serviceResponse.MostNumerousSpeciesQty).To(Equal(expectedQty))
				})

				It("Error - Some error with database", func() {

					expectedErr := errors.New("Any error")

					mockPetRepository.On("MostNumerousSpecies", ctx).Return(nil, expectedErr).Once()

					_, err := petService.MostNumerousSpecies(ctx)
					Expect(err).To(Equal(expectedErr))
				})
			})

			Context("Scenario - Database is empty", func() {
				It("Sucess - Return an error", func() {

					expectedErr := errors.New("species not found")

					mockPetRepository.On("MostNumerousSpecies", ctx).Return(nil, expectedErr).Once()

					serviceResponse, err := petService.MostNumerousSpecies(ctx)
					Expect(err).NotTo(BeNil())
					Expect(serviceResponse).To(BeNil())
				})
			})
		})

		Describe("AverageAge", func() {
			Context("Scenario - Database has some pets", func() {
				It("Success - Returns average age for the required species", func() {

					speciesName := "Cat"
					//expectedAge := 2.5
					birthday1, _ := time.Parse("2021-01-01", "2006-01-02")
					birthday2, _ := time.Parse("2022-01-01", "2006-01-02")

					mockPetRepository.On("GetAllPetsBySpecies", ctx, speciesName).Return([]domain.Pet{
						{Birthday: birthday1},
						{Birthday: birthday2},
					}, nil).Once()

					serviceResponse, err := petService.AverageAge(ctx, &speciesName)
					Expect(err).To(BeNil())
					//Expect(serviceResponse.AverageAge).To(BeNearTolerance(expectedAge, 0.1))
					Expect(serviceResponse.AverageAge).NotTo(BeZero())
					Expect(serviceResponse.AgeStandardDeviation).NotTo(BeZero())
				})

				It("Error - Error accesing db", func() {

					speciesName := "Cat"
					expectedErr := errors.New("Any error")

					mockPetRepository.On("GetAllPetsBySpecies", ctx, speciesName).Return(nil, expectedErr).Once()

					_, err := petService.AverageAge(ctx, &speciesName)
					Expect(err).To(Equal(expectedErr))
				})
			})
		})
	*/
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
