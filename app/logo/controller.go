package logo

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (controller Controller) Get(c echo.Context) error {
	url := "./assets/membership.png"

	file, err := os.Stat(url)
	if err != nil {
		return echo.ErrInternalServerError
	}

	modifiedTime := file.ModTime()
	etag := fmt.Sprintf("%x", md5.Sum([]byte(modifiedTime.String())))
	c.Response().Header().Set("ETag", etag)

	for _, tag := range c.Request().Header["If-None-Match"] {
		if tag == etag {
			return c.JSON(http.StatusNotModified, nil)
		}
	}

	return c.File(url)
}
