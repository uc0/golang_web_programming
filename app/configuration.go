package app

import (
	"comento_git_practice/app/logo"
	"comento_git_practice/app/membership"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

type Config struct {
	MembershipController membership.Controller
	LogoController       logo.Controller
}

func DefaultConfig() *Config {
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)
	return &Config{
		MembershipController: *controller,
		LogoController:       *logo.NewController(),
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

	e.GET("/memberships/:id", config.MembershipController.GetByID)
	e.GET("/memberships", config.MembershipController.GetMany)
	e.POST("/memberships", config.MembershipController.Create)
	e.PUT("/memberships", config.MembershipController.Update)
	e.DELETE("/memberships/:id", config.MembershipController.Delete)

	e.GET("/logo", config.LogoController.Get)

	return e
}
