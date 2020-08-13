package datasource_test

import (
	"encoding/json"
	"log"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/data/datasource"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/di"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/model"
)

var ds datasource.CustomerDataSource

func init() {
	ds = datasource.NewCustomerDataSource(di.ProvideDatabaseHelper())
}

func mockCustomer() model.Customer {
	jsonString := `
	{
		"id": 97,
		"code": "ACM0097",
		"name": "zhamppx"
	}`
	var data model.Customer
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		log.Println("Cannot Unmarshal")
	}
	return data
}