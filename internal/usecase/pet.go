package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/totoyk/trial-api-golang/internal/domain/model"
	"github.com/totoyk/trial-api-golang/internal/domain/repository"
)

type PetReceiver interface {
	// (GET /pets)
	ListPets(ctx echo.Context) (*[]model.Pet, error)

	// (GET /pets/{id})
	FindPetById(ctx echo.Context, id uint) (*model.Pet, error)
}

type PetInteractor struct {
	petRepository repository.PetRepository
}

func NewPetInteractor(petRepository repository.PetRepository) *PetInteractor {
	return &PetInteractor{
		petRepository: petRepository,
	}
}

func (i *PetInteractor) ListPets(ctx echo.Context) (*[]model.Pet, error) {
	pets, err := i.petRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return &pets, nil
}

func (i *PetInteractor) FindPetById(ctx echo.Context, id uint) (*model.Pet, error) {
	pet, err := i.petRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}
