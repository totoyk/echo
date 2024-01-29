package interfaces

import (
	"github.com/totoyk/trial-api-golang/internal/domain/repository"
	"github.com/totoyk/trial-api-golang/internal/usecase"
	"gorm.io/gorm"
)

type Handlers struct {
	PetHandler
}

func NewHandlers(dbconn *gorm.DB) *Handlers {
	// DI repository
	petRepository := repository.NewPetRepository(dbconn)

	// DI interfaces -> usecase
	return &Handlers{
		PetHandler: *NewPetHandler(usecase.NewPetInteractor(*petRepository)),
	}
}
