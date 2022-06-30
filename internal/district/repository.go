package district

import (
	"fmt"

	"github.com/anousoneFS/go-fiber-postgres-workshop/internal/province"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() (p []District, err error)
	Create(p *District) (id uint, err error)
	Update(id uint, p District) error
	GetByID(id uint) (District, error)
	Delete(i uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// table: Distrcit schema
type District struct {
	gorm.Model
	Name       string            `json:"name"`
	NameEn     string            `json:"name_en"`
	Province   province.Province // fk => table province
	ProvinceID uint
}

func (r repository) GetAll() (p []District, err error) {
	err = r.db.Find(&p).Error
	return p, err
}

func (r repository) Create(p *District) (id uint, err error) {
	err = r.db.Create(p).Error
	fmt.Printf("repo: %v\n", p)
	return p.ID, err
}

func (r repository) Update(id uint, p District) error {
	return r.db.Model(&District{}).Where("id = ?", id).Updates(p).Error
}

func (r repository) GetByID(id uint) (District, error) {
	var district District
	return district, r.db.Where("id=?", id).Find(&district).Error
}

func (r repository) Delete(id uint) error {
	return r.db.Delete(&District{}, id).Error
}
