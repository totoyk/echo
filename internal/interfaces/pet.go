package interfaces

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/totoyk/trial-api-golang/internal/interfaces/response"
	"github.com/totoyk/trial-api-golang/internal/usecase"
)

type PetHandler struct {
	UsecaseReceiver usecase.PetReceiver
}

func NewPetHandler(petReceiver usecase.PetReceiver) *PetHandler {
	return &PetHandler{
		UsecaseReceiver: petReceiver,
	}
}

// (GET /pets)
func (h *PetHandler) ListPets(ctx echo.Context) error {
	result, err := h.UsecaseReceiver.ListPets(ctx)
	if err != nil {
		return err
	}
	res := response.NewPets(result)
	return ctx.JSON(http.StatusOK, res)
}

// (GET /pets/{id})
func (h *PetHandler) FindPetById(ctx echo.Context, id uint) error {
	result, err := h.UsecaseReceiver.FindPetById(ctx, id)
	if err != nil {
		return err
	}
	res := response.NewPet(result)
	return ctx.JSON(http.StatusOK, res)
}
