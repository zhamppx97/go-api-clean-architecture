package repository

import (
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/data/datasource"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/model"
)

type CustomerRepository interface {
	Add(data *model.Customer) error
	GetAll() ([]model.Customer, error)
	Get(id int) (model.Customer, error)
}

type customerRepository struct {
	Ds datasource.CustomerDataSource
}

func NewCustomerRepository(ds datasource.CustomerDataSource) CustomerRepository {
	return &customerRepository{
		Ds: ds,
	}
}

func (repo *customerRepository) Add(data *model.Customer) error {
	return repo.Ds.Add(data)
}

func (repo *customerRepository) GetAll() ([]model.Customer, error) {
	return repo.Ds.GetAll()
}

func (repo *customerRepository) Get(id int) (model.Customer, error) {
	return repo.Ds.Get(id)
}