package authentication

import (
  "net/http"

  "github.com/gorilla/sessions"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/yinhr/todo-api/middleware/session"
  "github.com/yinhr/todo-api/util"
  log "github.com/sirupsen/logrus"
)

type (
  Config struct {
    Skipper middleware.Skipper
  }
  SkippablePair struct {
    Path string
    Method string
  }
)

var (
  SkippablePairs = []SkippablePair {
    SkippablePair{"/signup", http.MethodPost},
    SkippablePair{"/signin", http.MethodPost},
  }
  AuthenticationSkipper = func(c echo.Context) bool {
    for _, pair := range SkippablePairs {
      if pair.Path == c.Path() && pair.Method == c.Request().Method {
        return true
      }
    }
    return false
  }
  DefaultConfig = Config {
    Skipper: AuthenticationSkipper,
  }
)

func AuthenticationWithDefaultConfig() echo.MiddlewareFunc {
  config := DefaultConfig
  return AuthenticationWithConfig(config)
}

func AuthenticationWithConfig(config Config) echo.MiddlewareFunc {
  if config.Skipper == nil {
    config.Skipper = middleware.DefaultSkipper
  }
  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      if config.Skipper(c) {
        return next(c)
      }
      store, _ := c.Get(session.Key).(sessions.Store)
      sess, _ := store.Get(c.Request(), session.Name)
      if _, ok := sess.Values[session.UID]; ok {
        sess.Options = &sessions.Options {
          Path: "/",
          MaxAge: 3600 * 24,
          HttpOnly: true,
          Secure: util.Production(),
        }
        err := sessions.Save(c.Request(), c.Response())
        if err != nil {
          log.Error(err)
          return c.JSON(http.StatusInternalServerError, util.NewInternalServerError())
        }
        return next(c)
      } else {
        return c.JSON(http.StatusUnauthorized, util.NewUnauthorizedError("Unauthorized"))
      }
    }
  }
}
