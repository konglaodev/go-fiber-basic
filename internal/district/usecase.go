package district

import "fmt"

type Usecase interface {
	GetAll() (p []District, err error)
	Create(p DistrictRequest) (id uint, err error)
	Update(id uint, p DistrictRequest) error
	GetByID(id uint) (District, error)
	Delete(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo: repo}
}

func (u usecase) Create(p DistrictRequest) (id uint, err error) {
	District := District{Name: p.Name, NameEn: p.NameEn, ProvinceID: p.ProvinceID}
	id, err = u.repo.Create(&District)
	if err != nil {
		fmt.Printf("usecase.GetAll():%v\n", err)
		return 0, err
	}
	return
}

func (u usecase) GetAll() (p []District, err error) {
	i, err := u.repo.GetAll()
	if err != nil {
		// log
		fmt.Printf("usecase.GetAll():%v\n", err)
		return []District{}, err
	}
	return i, err
}

func (u usecase) GetByID(id uint) (District, error) {
	i, err := u.repo.GetByID(id)
	if err != nil {
		// log
		fmt.Printf("usecase.GetByID():%v\n", err)
		return District{}, err
	}
	return i, err
}

func (u usecase) Update(id uint, p DistrictRequest) error {
	newDistrict := District{}
	if p.Name != "" {
		newDistrict.Name = p.Name
	}
	if p.NameEn != "" {
		newDistrict.NameEn = p.NameEn
	}
	if err := u.repo.Update(id, newDistrict); err != nil {
		fmt.Printf("usecase.Update(): %v\n", err)
		return err
	}
	return nil
}

func (u usecase) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		fmt.Printf("usecase.Delete(): %v\n", err)
		return err
	}
	return nil
}
