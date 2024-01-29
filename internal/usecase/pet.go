package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/totoyk/trial-api-golang/internal/domain/model"
	"github.com/totoyk/trial-api-golang/internal/domain/repository"
	"github.com/totoyk/trial-api-golang/internal/oas"
	"github.com/totoyk/trial-api-golang/internal/pkg"
)

type PetReceiver interface {
	// (GET /pets)
	ListPets(ctx echo.Context) (*[]model.Pet, error)

	// (GET /pets/{id})
	FindPetById(ctx echo.Context, id uint) (*model.Pet, error)

	// (POST /pets)
	CreatePets(ctx echo.Context, params oas.PetRequestBody) error

	// (PUT /pets/{id})
	UpdatePet(ctx echo.Context, id uint, params oas.PetRequestBody) error

	// (DELETE /pets/{id})
	DeletePet(ctx echo.Context, id uint) error
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

func (i *PetInteractor) CreatePets(ctx echo.Context, params oas.PetRequestBody) error {
	pet := model.Pet{
		Name: params.Name,
		Tag:  pkg.NullString(params.Tag),
	}
	err := i.petRepository.Create(&pet)
	if err != nil {
		return err
	}
	return nil
}

func (i *PetInteractor) UpdatePet(ctx echo.Context, id uint, params oas.PetRequestBody) error {
	pet, err := i.petRepository.FindById(id)
	if err != nil {
		return err
	}
	pet.Name = params.Name
	pet.Tag = pkg.NullString(params.Tag)
	err = i.petRepository.Update(&pet)
	if err != nil {
		return err
	}
	return nil
}

func (i *PetInteractor) DeletePet(ctx echo.Context, id uint) error {
	pet, err := i.petRepository.FindById(id)
	if err != nil {
		return err
	}
	err = i.petRepository.Delete(&pet)
	if err != nil {
		return err
	}
	return nil
}
