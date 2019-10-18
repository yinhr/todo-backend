package delivery

import (
  "net/http"

  "github.com/gorilla/sessions"
  "github.com/labstack/echo/v4"
  "github.com/yinhr/todo-api/middleware/session"
  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/usecase/signup"
  "github.com/yinhr/todo-api/util"
  log "github.com/sirupsen/logrus"
)

type (
  SignupHandler struct {
    SignupUsecase signup.Usecase
  }
)

func (h *SignupHandler) Post(c echo.Context) error {
  u := new(model.User)
  if err := c.Bind(u); err != nil {
    log.Error(err)
    return c.JSON(http.StatusBadRequest, util.NewBadRequestError("400: BadRequest"))
  }
  err := h.SignupUsecase.Store(util.GetContext(c), u)
  if err != nil {
    return c.JSON(http.StatusBadRequest, err)
  }
  sess := session.NewLoginSession(c)
  sess.Values[session.UID] = u.ID
  err = sessions.Save(c.Request(), c.Response())
  if err != nil {
    log.Error(err)
    return c.JSON(http.StatusInternalServerError, util.NewInternalServerError())
  }
  return c.JSON(http.StatusOK, nil)
}

func (h *SignupHandler) Index(c echo.Context) error {
  u, _ := h.SignupUsecase.Fetch(util.GetContext(c))
  return c.JSON(http.StatusOK, u)
}
