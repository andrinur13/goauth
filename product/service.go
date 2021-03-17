package product

import "errors"

type Service interface {
	SaveProduct(product ProductInput) (Product, error)
	GetProducts() ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveProduct(input ProductInput) (Product, error) {

	newProduct := Product{
		Name:  input.Name,
		Price: input.Price,
	}

	dataProduct, err := s.repository.Save(newProduct)

	if err != nil {
		return dataProduct, errors.New("Failed store Product!")
	}

	return dataProduct, nil

}

func (s *service) GetProducts() ([]Product, error) {
	data, err := s.repository.GetAll()

	if err != nil {
		return data, err
	}

	if len(data) < 1 {
		return data, errors.New("Data Not Found!")
	}

	return data, nil
}
