package membership

import (
	"github.com/labstack/echo"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller Controller) Create(c echo.Context) error {
	req := new(CreateRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := controller.service.Create(*req)
	return c.JSON(res.Code, res)
}

func (controller Controller) GetByID(c echo.Context) error {
	res := controller.service.GetByID(c.Param("id"))

	return c.JSON(http.StatusOK, res)
}

func (controller Controller) Update(c echo.Context) error {
	req := new(UpdateRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := controller.service.Update(*req)
	return c.JSON(res.Code, res)
}

func (controller Controller) Delete(c echo.Context) error {
	res := controller.service.Delete(c.Param("id"))

	return c.JSON(res.Code, res)
}
