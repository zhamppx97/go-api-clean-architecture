package main

func main() {
	e := echo.New()
	// Routes
	route.NewCustomerRoute(di.ProvideCustomerHandler()).Initial(e)
	// Listener
	e.Logger.Fatal(e.Start(":1323"))
}