package handler

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/domain"
	"github.com/zhamppx97/go-api-clean-architecture/api/v1/customer/model"
)

// CustomerHandler : interface
type CustomerHandler interface {
	Add(c echo.Context) error
	GetAll(c echo.Context) error
	Get(c echo.Context) error
}

type customerHandler struct {
	UseCase domain.CustomerUseCase
}

// NewCustomerHandler : new customer handler
func NewCustomerHandler(useCase domain.CustomerUseCase) CustomerHandler {
	return &customerHandler{
		UseCase: useCase,
	}
}

func (h *customerHandler) Add(c echo.Context) error {
	var customer model.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Data not found!",
		})
	}
	if err := h.UseCase.Add(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}
	return c.JSON(http.StatusCreated, customer)
}

func (h *customerHandler) GetAll(c echo.Context) error {
	customers, _ := h.UseCase.GetAll()
	return c.JSON(http.StatusOK, customers)
}

func (h *customerHandler) Get(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		idInt, iErr := strconv.Atoi(id)
		if idInt > 0 && iErr == nil {
			customer, err := h.UseCase.Get(idInt)
			if err == nil {
				return c.JSON(http.StatusOK, customer)
			}
			return c.JSON(http.StatusNotFound, err)
		}
	}
	return c.JSON(http.StatusBadGateway, echo.Map{
		"message": "Required parameter id",
	})
}