package domain

import (
	"fmt"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/data/repository"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/model"
)

// CustomerUseCase : interface
type CustomerUseCase interface {
	Add(data *model.Customer) error
	GetAll() ([]model.Customer, error)
	Get(id int) (model.Customer, error)
}

type customerUseCase struct {
	Repo repository.CustomerRepository
}

// NewCustomerUseCase : new customer use case
func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		Repo: repo,
	}
}

func (uc *customerUseCase) Add(data *model.Customer) error {
	if err := uc.Repo.Add(data); err == nil {
		cus, err := uc.Get(data.ID)
		*data = cus
		return err
	}
	return fmt.Errorf("Not found")
}

func (uc *customerUseCase) GetAll() ([]model.Customer, error) {
	return uc.Repo.GetAll()
}

func (uc *customerUseCase) Get(id int) (model.Customer, error) {
	return uc.Repo.Get(id)
}