package main

import (
	"github.com/labstack/echo"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/di"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/gateway/route"
)

func main() {
	e := echo.New()
	// Routes
	route.NewCustomerRoute(di.ProvideCustomerHandler()).Initial(e)
	// Listener
	e.Logger.Fatal(e.Start(":7000"))
}