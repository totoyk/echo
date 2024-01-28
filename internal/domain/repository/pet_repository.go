package repository

import (
	"github.com/totoyk/trial-api-golang/internal/domain/model"
	"gorm.io/gorm"
)

type PetRepository struct {
	Conn *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepository {
	return &PetRepository{
		Conn: db,
	}
}

func (r *PetRepository) FindAll() (pets []model.Pet, err error) {
	err = r.Conn.Find(&pets).Error
	return
}

func (r *PetRepository) FindById(id int64) (pet model.Pet, err error) {
	err = r.Conn.First(&pet, id).Error
	return
}

func (r *PetRepository) Create(pet *model.Pet) (err error) {
	err = r.Conn.Create(pet).Error
	return
}

func (r *PetRepository) Update(pet *model.Pet) (err error) {
	err = r.Conn.Save(pet).Error
	return
}

func (r *PetRepository) Delete(pet *model.Pet) (err error) {
	err = r.Conn.Delete(pet).Error
	return
}
