package route

import (
	"github.com/labstack/echo"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/gateway/handler"
)

// CustomerRoute : interface
type CustomerRoute interface {
	Initial(e *echo.Echo)
}

type customerRoute struct {
	Handle handler.CustomerHandler
}

// NewCustomerRoute : new customer route
func NewCustomerRoute(handler handler.CustomerHandler) CustomerRoute {
	return &customerRoute{
		Handle: handler,
	}
}

func (r *customerRoute) Initial(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("/customer", r.Handle.GetAll)
	v1.POST("/customer", r.Handle.Add)
	v1.GET("/customer/:id", r.Handle.Get)
}