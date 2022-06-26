package app

import (
	"comento_git_practice/app/logo"
	"comento_git_practice/app/membership"
	"comento_git_practice/app/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

type Config struct {
	MembershipController membership.Controller
	LogoController       logo.Controller
	UserController       user.Controller
	UserMiddleware       user.Middleware
}

func DefaultConfig() *Config {
	data := map[string]membership.Membership{}

	membershipService := membership.NewService(*membership.NewRepository(data))
	membershipController := membership.NewController(*membershipService)

	userService := user.NewService(user.DefaultSecret)

	return &Config{
		MembershipController: *membershipController,
		LogoController:       *logo.NewController(),
		UserController:       *user.NewController(*userService),
		UserMiddleware:       *user.NewMiddleware(*membership.NewRepository(data)),
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

	membershipController := config.MembershipController

	userController := config.UserController
	userMiddleware := config.UserMiddleware

	logoController := config.LogoController

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{Claims: &user.Claims{}, SigningKey: user.DefaultSecret})

	e.POST("/login", userController.Login)
	e.GET("/memberships/:id", config.MembershipController.GetByID)
	e.GET("/memberships", membershipController.GetMany)
	e.POST("/memberships", membershipController.Create)
	e.PUT("/memberships", membershipController.Update)
	e.DELETE("/memberships/:id", membershipController.Delete)
	e.GET("/logo", logoController.Get)

	return e
}
