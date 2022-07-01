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

func (m Middleware) ValidateOnlyMemberByParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := m.validateAuth(c, false, false)
		if err != nil {
			return err
		}

		return next(c)
	}
}

func (m Middleware) ValidateMemberByParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := m.validateAuth(c, false, true)
		if err != nil {
			m.ValidateAdmin(next)
		}

		return next(c)
	}
}

func (m Middleware) ValidateOnlyMemberByBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := m.validateAuth(c, true, false)
		if err != nil {
			return err
		}

		return next(c)
	}
}

func (m Middleware) ValidateMemberByBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := m.validateAuth(c, true, true)
		if err != nil {
			m.ValidateAdmin(next)
		}

		return next(c)
	}
}

func (m Middleware) validateAuth(c echo.Context, hasBody bool, canAdmin bool) error {
	isUnauthorized := false
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims)
	mem, _ := m.membershipRepository.GetByUserName(claims.Name)

	if hasBody {
		body := make(map[string]string)
		if err := c.Bind(body); err != nil {
			isUnauthorized = true
		}
		if body["id"] != mem.ID {
			isUnauthorized = true
		}

		if isUnauthorized && !canAdmin {
			return echo.ErrUnauthorized
		}

		return nil
	}

	if c.Param("id") != mem.ID {
		isUnauthorized = true
	}

	if isUnauthorized && !canAdmin {
		return echo.ErrUnauthorized
	}

	return nil
}
