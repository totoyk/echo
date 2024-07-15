package interfaces

import (
	"github.com/totoyk/sample-backendapp-golang-on-echo/internal/domain/repository"
	"github.com/totoyk/sample-backendapp-golang-on-echo/internal/usecase"
	"gorm.io/gorm"
)

type Handlers struct {
	PetHandler
}

func NewHandlers(dbconn *gorm.DB) *Handlers {
	// DI repository <- infrastructure
	petRepository := repository.NewPetRepository(dbconn)

	// DI interfaces -> usecase -> repository
	return &Handlers{
		PetHandler: *NewPetHandler(usecase.NewPetInteractor(*petRepository)),
	}
}
