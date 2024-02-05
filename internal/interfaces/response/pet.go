package response

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/totoyk/trial-api-golang/internal/domain/model"
	"github.com/totoyk/trial-api-golang/internal/oas"
	"github.com/totoyk/trial-api-golang/internal/pkg"
)

func NewPet(ctx echo.Context, m *model.Pet) oas.Pet {
	return oas.Pet{
		Id:   m.ID,
		Name: m.Name,
		Tag: func() *string {
			if m.Tag.Valid {
				return &m.Tag.String
			}
			return nil
		}(),
		DateOfBirth: func() *time.Time {
			if m.DateOfBirth.Valid {
				// TODO: use ctx to get the timezone
				return pkg.SetLocation(&m.DateOfBirth.Time, "Asia/Tokyo")
			}
			return nil
		}(),
	}
}

func NewPets(ctx echo.Context, ms *[]model.Pet) []oas.Pet {
	var pets []oas.Pet
	for _, m := range *ms {
		pets = append(pets, NewPet(ctx, &m))
	}
	return pets
}
