package delivery

import (
  "net/http"
  "strconv"

  "github.com/labstack/echo/v4"
  "github.com/yinhr/todo-api/middleware/session"
  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/usecase/todo"
  "github.com/yinhr/todo-api/util"
  log "github.com/sirupsen/logrus"
)

type (
  TodoHandler struct {
    TodoUsecase todo.Usecase
  }
)

func (h *TodoHandler) FindBy(c echo.Context) error {
  id := c.QueryParam("id")
  todo, err := h.TodoUsecase.FindBy(util.GetContext(c), id)
  if err != nil {
    he, _ := err.(*echo.HTTPError)
    return c.JSON(he.Code, he)
  }
  return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Index(c echo.Context) error {
  params := new(util.QueryParams)
  sess, err := session.Get(session.Name, c)
  if err != nil {
    log.Error(err)
    return c.JSON(http.StatusInternalServerError, util.NewInternalServerError())
  }
  params.UserID, _ = sess.Values[session.UID].(string)
  params.Cursor, _ = strconv.Atoi(c.QueryParam("cursor"))
  params.OrderBy = c.QueryParam("orderBy")
  params.Direction = c.QueryParam("direction")
  todos, err := h.TodoUsecase.Fetch(util.GetContext(c), params)
  if err != nil {
    return c.JSON(http.StatusBadRequest, util.NewBadRequestError("400: BadRequest"))
  }
  return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) Post(c echo.Context) error {
  t := new(model.Todo)
  if err := c.Bind(t); err != nil {
    log.Error(err)
    return c.JSON(http.StatusBadRequest, util.NewBadRequestError("400: BadRequest"))
  }
  sess, err := session.Get(session.Name, c)
  if err != nil {
    log.Error(err)
    return c.JSON(http.StatusInternalServerError, util.NewInternalServerError())
  }
  t.UserID, _ = sess.Values[session.UID].(string)
  if err := h.TodoUsecase.Create(util.GetContext(c), t); err != nil {
    he, _ := err.(*echo.HTTPError)
    return c.JSON(he.Code, he)
  }
  return c.JSON(http.StatusOK, nil)
}

func (h *TodoHandler) PatchTodo(c echo.Context) error {
  t := new(model.Todo)
  if err := c.Bind(t); err != nil {
    log.Error(err)
    return c.JSON(http.StatusBadRequest, util.NewBadRequestError("400: BadRequest"))
  }
  if err := h.TodoUsecase.PatchTodoDone(util.GetContext(c), t); err != nil {
    he, _ := err.(*echo.HTTPError)
    return c.JSON(he.Code, he)
  }
  return c.JSON(http.StatusOK, nil)
}

func (h *TodoHandler) PatchTask(c echo.Context) error {
  t := new(model.Task)
  if err := c.Bind(t); err != nil {
    log.Error(err)
    return c.JSON(http.StatusBadRequest, util.NewBadRequestError("400: BadRequest"))
  }
  if err := h.TodoUsecase.PatchTaskDone(util.GetContext(c), t); err != nil {
    he, _ := err.(*echo.HTTPError)
    return c.JSON(he.Code, he)
  }
  return c.JSON(http.StatusOK, nil)
}

func (h *TodoHandler) Put(c echo.Context) error {
  t := new(model.Todo)
  if err := c.Bind(t); err != nil {
    log.Error(err)
    return c.JSON(http.StatusBadRequest, util.NewBadRequestError("400: BadRequest"))
  }
  if err := h.TodoUsecase.Put(util.GetContext(c), t); err != nil {
    he, _ := err.(*echo.HTTPError)
    return c.JSON(he.Code, he)
  }
  return c.JSON(http.StatusOK, nil)
}

func (h *TodoHandler) Destroy(c echo.Context) error {
  id := c.QueryParam("id")
  err := h.TodoUsecase.Delete(util.GetContext(c), id)
  if err != nil {
    he, _ := err.(*echo.HTTPError)
    return c.JSON(he.Code, he)
  }
  return c.JSON(http.StatusOK, nil)
}
