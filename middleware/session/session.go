package session

import (
  "github.com/gorilla/context"
  "github.com/gorilla/sessions"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  log "github.com/sirupsen/logrus"
  "github.com/yinhr/todo-api/util"
)

type (
  Config struct {
    Skipper middleware.Skipper
    Store sessions.Store
  }
)

var (
  DefaultConfig = Config {
    Skipper: middleware.DefaultSkipper,
  }
)

const (
  Key = "_session_store"
  Name = "sessionid"
  UID = "uid"
)

func SessionWithDefaultConfig(store sessions.Store) echo.MiddlewareFunc {
  config := DefaultConfig
  config.Store = store
  return SessionWithConfig(config)
}

func SessionWithConfig(config Config) echo.MiddlewareFunc {
  if config.Skipper == nil {
    config.Skipper = middleware.DefaultSkipper
  }
  if config.Store == nil {
    panic("echo: session middleware requires store")
  }
  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      if config.Skipper(c) {
        return next(c)
      }
      defer context.Clear(c.Request())
      c.Set(Key, config.Store)
      return next(c)
    }
  }
}

func New(name string, c echo.Context) (*sessions.Session, error) {
  store, _ := c.Get(Key).(sessions.Store)
  return store.New(c.Request(), name)
}

func Get(name string, c echo.Context) (*sessions.Session, error) {
  store, _ := c.Get(Key).(sessions.Store)
  return store.Get(c.Request(), name)
}

func Save(s *sessions.Session, c echo.Context) error {
  store, _ := c.Get(Key).(sessions.Store)
  return store.Save(c.Request(), c.Response(), s)
}

func Delete(name string, c echo.Context) error {
  store, _ := c.Get(Key).(sessions.Store)
  sess, err := store.Get(c.Request(), Name)
  if err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  sess.Options = &sessions.Options {
    MaxAge: -1,
    Path: "/",
  }
  err = store.Save(c.Request(), c.Response(), sess)
  if err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  return nil
}

func NewLoginSession(c echo.Context) (*sessions.Session) {
  store, _ := c.Get(Key).(sessions.Store)
  sess, _ := store.Get(c.Request(), Name)
  sess.Options = &sessions.Options { 
    Path: "/",
    MaxAge: 3600 * 24,
    HttpOnly: true,
    Secure: util.Production(),
  }
  return sess
}
