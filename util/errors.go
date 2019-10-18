package util

import (
  "net/http"

  "github.com/go-sql-driver/mysql"
  "github.com/labstack/echo/v4"
)

func NewBadRequestError(message string) error {
  return echo.NewHTTPError(http.StatusBadRequest, message)
}

func NewUnauthorizedError(message string) error {
  return echo.NewHTTPError(http.StatusUnauthorized, message)
}

func NewNotFoundError(message string) error {
  return echo.NewHTTPError(http.StatusNotFound, message)
}

func NewTimeoutError(message string) error {
  return echo.NewHTTPError(http.StatusRequestTimeout, message)
}

func NewInternalServerError() error {
  return echo.NewHTTPError(http.StatusInternalServerError, "500 Internal Server Error")
}

func NewDuplicateEntryError() error {
  return echo.NewHTTPError(http.StatusBadRequest, "既に登録されています")
}

func SqlErrorHandler(e error) error {
  if err, ok := e.(*mysql.MySQLError); ok {
    switch (*err).Number {
    case 1062: return NewDuplicateEntryError()
    default: return NewInternalServerError()
    }
  } else {
    return NewInternalServerError()
  }
}
