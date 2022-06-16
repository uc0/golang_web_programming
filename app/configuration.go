package app

import (
	"comento_git_practice/app/membership"
	"github.com/labstack/echo"
	"net/http"
)

type Config struct {
	MembershipApplication membership.Application
}

func DefaultConfig() *Config {
	data := map[string]membership.Membership{}
	application := membership.NewApplication(*membership.NewRepository(data))
	return &Config{
		MembershipApplication: *application,
	}
}

func NewEcho(config Config) *echo.Echo {
	e := echo.New()

	e.GET("/memberships", func(c echo.Context) error {
		// config.MembershipApplication.Get()
		return c.JSON(http.StatusOK, "hello")
	})

	return e
}
