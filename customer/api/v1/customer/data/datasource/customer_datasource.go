package datasource

import (
	"fmt"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/model"
)

func GetDatabaseMock() DatabaseHelper {
	return DatabaseHelper{
		Store: make(map[int]model.Customer),
	}
}

type DatabaseHelper struct {
	Store map[int]model.Customer
}

type CustomerDataSource interface {
	Add(data *model.Customer) error
	GetAll() ([]model.Customer, error)
	Get(id int) (model.Customer, error)
}

type customerDataSource struct {
	Db DatabaseHelper
}

func NewCustomerDataSource(db DatabaseHelper) CustomerDataSource {
	return &customerDataSource{
		Db: db,
	}
}

func (repo *customerDataSource) Add(data *model.Customer) error {
	if data.ID > 0 {
		repo.Db.Store[data.ID] = *data
		return nil
	}
	return fmt.Errorf("Cannot add data: %s", "is empty")
}

func (repo *customerDataSource) GetAll() ([]model.Customer, error) {
	data := []model.Customer{}
	for _, value := range repo.Db.Store {
		data = append(data, value)
	}
	return data, nil
}

func (repo *customerDataSource) Get(id int) (model.Customer, error) {
	data := repo.Db.Store[id]
	if data.ID == id {
		return repo.Db.Store[id], nil
	}
	return model.Customer{}, fmt.Errorf("Not found id: %d", id)
}