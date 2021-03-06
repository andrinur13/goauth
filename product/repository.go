package product

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(product Product) (Product, error)
	GetAll() ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product

	err := r.db.Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}
