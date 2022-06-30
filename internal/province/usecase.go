package province

import "fmt"

type Usecase interface {
	GetAll() (p []Province, err error)
	Create(p ProvinceRequest) (id uint, err error)
	Update(id uint, p ProvinceRequest) error
	GetByID(id uint) (Province, error)
	Delete(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo: repo}
}

func (u usecase) Create(p ProvinceRequest) (id uint, err error) {
	province := Province{Name: p.Name, NameEn: p.NameEn}
	id, err = u.repo.Create(&province)
	if err != nil {
		fmt.Printf("usecase.GetAll():%v\n", err)
		return 0, err
	}
	return
}

func (u usecase) GetAll() (p []Province, err error) {
	i, err := u.repo.GetAll()
	if err != nil {
		// log
		fmt.Printf("usecase.GetAll():%v\n", err)
		return []Province{}, err
	}
	return i, err
}

func (u usecase) GetByID(id uint) (Province, error) {
	i, err := u.repo.GetByID(id)
	if err != nil {
		// log
		fmt.Printf("usecase.GetByID():%v\n", err)
		return Province{}, err
	}
	return i, err
}

func (u usecase) Update(id uint, p ProvinceRequest) error {
	newProvince := Province{}
	if p.Name != "" {
		newProvince.Name = p.Name
	}
	if p.NameEn != "" {
		newProvince.NameEn = p.NameEn
	}
	if err := u.repo.Update(id, newProvince); err != nil {
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
