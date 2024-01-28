package ui

import (
	"github.com/totoyk/trial-api-golang/internal/domain/repository"
	"github.com/totoyk/trial-api-golang/internal/usecase"
	"gorm.io/gorm"
)

type Handlers struct {
	PetHandler
}

func NewHandlers(dbconn *gorm.DB) *Handlers {
	repositories := repository.NewRepositories(dbconn)

	return &Handlers{
		PetHandler: PetHandler{
			UsecaseReceiver: usecase.NewPetInteractor(*repositories.PetRepository),
		},
	}
}
