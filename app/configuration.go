package app

import (
	"comento_git_practice/app/membership"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

type Config struct {
	Controller membership.Controller
}

func DefaultConfig() *Config {
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)
	return &Config{
		Controller: *controller,
	}
}

func NewEcho(config Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("Request URI: ", c.Request().RequestURI)
		log.Println("Request Http Method: ", c.Request().Method)
		if string(reqBody) != "" {
			log.Println("Request Body: ", string(reqBody))
		}
		log.Println("Response Http Status Code: ", c.Response().Status)
		if string(resBody) != "" {
			log.Println("Response Body: ", string(resBody))
		}
	}))

	controller := config.Controller

	e.GET("/memberships/:id", controller.GetByID)
	e.GET("/memberships", controller.GetMany)
	e.POST("/memberships", controller.Create)
	e.PUT("/memberships", controller.Update)
	e.DELETE("/memberships/:id", controller.Delete)

	return e
}
