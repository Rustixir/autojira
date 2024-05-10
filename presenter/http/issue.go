package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"smart/application/service"
)

type Jason map[string]interface{}

func Bootstrap(svc service.IssueService) {
	srv := echo.New()
	srv.Use(middleware.Logger())
	srv.Use(middleware.Recover())
	srv.POST("/", func(c echo.Context) error {
		req := new(struct {
			Message string
		})

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, Jason{"error": "invalid request"})
		}

		if err := svc.Create(c.Request().Context(), req.Message); err != nil {
			return c.JSON(http.StatusInternalServerError, Jason{"error": "internal server error"})
		}

		return c.NoContent(http.StatusOK)
	})
	srv.Start(os.Getenv("ADDRESS"))
}
