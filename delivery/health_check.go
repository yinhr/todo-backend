package delivery

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

type (
  HealthCheckHandler struct {}
)

func (h *HealthCheckHandler) Get(c echo.Context) error {
  return c.JSON(http.StatusOK, nil)
}
