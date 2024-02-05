package response

import (
	"time"

	"github.com/totoyk/trial-api-golang/internal/domain/model"
	"github.com/totoyk/trial-api-golang/internal/oas"
)

func NewPet(m *model.Pet) oas.Pet {
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
				return &m.DateOfBirth.Time
			}
			return nil
		}(),
	}
}

func NewPets(ms *[]model.Pet) []oas.Pet {
	var pets []oas.Pet
	for _, m := range *ms {
		pets = append(pets, NewPet(&m))
	}
	return pets
}
