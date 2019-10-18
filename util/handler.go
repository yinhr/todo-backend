package util

import (
  "context"
  "io"
  "net/http"
  "testing"

  "github.com/labstack/echo/v4"
)

func GetContext(c echo.Context) context.Context {
  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }
  return ctx
}

func NewHttpRequest(t *testing.T, method, path string, body io.Reader) *http.Request {
  t.Helper()
  req, err := http.NewRequest(method, path, body)
  if err != nil {
    t.Fatalf("expected no error, but got:'%s'", err)
  }
  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
  return req
}
