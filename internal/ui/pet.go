package ui

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/totoyk/trial-api-golang/internal/ui/interfaces"
	"github.com/totoyk/trial-api-golang/internal/usecase"
)

type PetHandler struct {
	UsecaseReceiver usecase.PetReceiver
}

func NewPetHandler() *PetHandler {
	return &PetHandler{}
}

// (GET /pets)
func (h *PetHandler) ListPets(ctx echo.Context) error {
	pets, err := h.UsecaseReceiver.ListPets(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, interfaces.NewPets(pets))
}

// (GET /pets/{id})
func (h *PetHandler) FindPetById(ctx echo.Context, id uint) error {
	pet, err := h.UsecaseReceiver.FindPetById(ctx, id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, interfaces.NewPet(pet))
}
