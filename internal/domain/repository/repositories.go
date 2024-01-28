package repository

import "gorm.io/gorm"

type Repositories struct {
	PetRepository *PetRepository
}

func NewRepositories(dbconn *gorm.DB) *Repositories {
	return &Repositories{
		PetRepository: NewPetRepository(dbconn),
	}
}
