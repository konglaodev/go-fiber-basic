package province

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() (p []Province, err error)
	Create(p *Province) (id uint, err error)
	Update(id uint, p Province) error
	GetByID(id uint) (Province, error)
	Delete(i uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// table: Province schema
type Province struct {
	gorm.Model
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

func (r repository) GetAll() (p []Province, err error) {
	err = r.db.Find(&p).Error
	return p, err
}

func (r repository) Create(p *Province) (id uint, err error) {
	err = r.db.Create(p).Error
	return p.ID, err
}

func (r repository) Update(id uint, p Province) error {
	return r.db.Model(&Province{}).Where("id = ?", id).Updates(p).Error
}

func (r repository) GetByID(id uint) (Province, error) {
	var province Province
	return province, r.db.Where("id=?", id).Find(&province).Error
}

func (r repository) Delete(id uint) error {
	return r.db.Delete(&Province{}, id).Error
}
