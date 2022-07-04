package user

import (
	"comento_git_practice/app/membership"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Middleware struct {
	membershipRepository membership.Repository
}

func NewMiddleware(membershipRepository membership.Repository) *Middleware {
	return &Middleware{membershipRepository: membershipRepository}
}

func (m Middleware) ValidateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		if !claims.IsAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func (m Middleware) ValidateMemberByParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		mem, _ := m.membershipRepository.GetByUserName(claims.Name)
		if c.Param("id") != mem.ID {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func (m Middleware) ValidateMemberOrAdminByParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		mem, _ := m.membershipRepository.GetByUserName(claims.Name)
		if c.Param("id") != mem.ID {
			m.ValidateAdmin(next)
		}

		return next(c)
	}
}

func (m Middleware) ValidateMemberByBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		mem, _ := m.membershipRepository.GetByUserName(claims.Name)
		body := make(map[string]string)
		if err := c.Bind(body); err != nil {
			return echo.ErrUnauthorized
		}
		if body["id"] != mem.ID {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
