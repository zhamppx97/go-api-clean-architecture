package di

import (
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/data/datasource"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/data/repository"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/domain"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/gateway/handler"
)

// ProvideDatabaseHelper : provide database helper
func ProvideDatabaseHelper() datasource.DatabaseHelper {
	return datasource.GetDatabaseMock()
}

// ProvideCustomerDataSource : provide datasource
func ProvideCustomerDataSource() datasource.CustomerDataSource {
	return datasource.NewCustomerDataSource(ProvideDatabaseHelper())
}

// ProvideCustomerRepository : provide repository
func ProvideCustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(ProvideCustomerDataSource())
}

// ProvideCustomerUseCase : provide usecase
func ProvideCustomerUseCase() domain.CustomerUseCase {
	return domain.NewCustomerUseCase(ProvideCustomerRepository())
}

// ProvideCustomerHandler : provide handler
func ProvideCustomerHandler() handler.CustomerHandler {
	return handler.NewCustomerHandler(ProvideCustomerUseCase())
}