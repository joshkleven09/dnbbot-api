package hello

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() (Hellos, error) {
	hellos := make([]*Hello, 0)
	if err := r.db.Find(&hellos).Error; err != nil {
		return nil, err
	}

	return hellos, nil
}

func (r *Repository) Find() (*Hello, error) {
	hello := &Hello{}
	if err := r.db.First(&hello).Error; err != nil {
		return nil, err
	}

	return hello, nil
}
