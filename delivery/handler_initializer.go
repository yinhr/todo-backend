package delivery

import (
  "net/http"

  "github.com/gorilla/sessions"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/yinhr/todo-api/middleware/session"
  "github.com/yinhr/todo-api/middleware/authentication"
  "github.com/yinhr/todo-api/usecase"
  "github.com/yinhr/todo-api/usecase/signup"
  "github.com/yinhr/todo-api/usecase/signin"
  "github.com/yinhr/todo-api/usecase/todo"
)

type (
  Handler int
  Handlers map[Handler]interface{}
)

const (
  SignupHdlr Handler = iota
  SigninHdlr
  TodoHdlr
  HealthCheckHdlr
)

func GetInitializedHandlers(ucases *usecase.UCases) Handlers {
  handlers := make(Handlers)

  signupUsecase, _ := (*ucases)[usecase.SignupUCase].(*signup.SignupUsecase) 
  handlers[SignupHdlr] = &SignupHandler{signupUsecase} 

  signinUsecase, _ := (*ucases)[usecase.SigninUCase].(*signin.SigninUsecase) 
  handlers[SigninHdlr] = &SigninHandler{signinUsecase} 

  todoUsecase, _ := (*ucases)[usecase.TodoUCase].(*todo.TodoUsecase) 
  handlers[TodoHdlr] = &TodoHandler{todoUsecase} 

  handlers[HealthCheckHdlr] = &HealthCheckHandler{}
  return handlers
}

func ConfigureHandler(e *echo.Echo, handlers *Handlers) {
  signupHandler, _ := (*handlers)[SignupHdlr].(*SignupHandler) 
  e.POST("/signup", signupHandler.Post)
  e.GET("/signup", signupHandler.Index)

  signinHandler, _ := (*handlers)[SigninHdlr].(*SigninHandler) 
  e.GET("/signin", signinHandler.Get)
  e.POST("/signin", signinHandler.Post)
  e.DELETE("/signin", signinHandler.Delete)

  todoHandler, _ := (*handlers)[TodoHdlr].(*TodoHandler) 
  e.GET("/todo", todoHandler.Index)
  e.GET("/todo/edit", todoHandler.FindBy)
  e.POST("/todo", todoHandler.Post)
  e.PUT("/todo", todoHandler.Put)
  e.PATCH("/todo", todoHandler.PatchTodo)
  e.PATCH("/task", todoHandler.PatchTask)
  e.DELETE("/todo", todoHandler.Destroy)

  healthCheckHandler, _ := (*handlers)[HealthCheckHdlr].(*HealthCheckHandler)
  e.GET("/healthcheck", healthCheckHandler.Get)
}

func ConfiguerMiddleware(e *echo.Echo, store sessions.Store) {
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.Use(sessionWithConfig(store))
  e.Use(corsWithConfig())
  e.Use(authentication.AuthenticationWithDefaultConfig())
}

func sessionWithConfig(store sessions.Store) echo.MiddlewareFunc {
  return session.SessionWithConfig(session.Config {
    Store: store,
  })
}

func corsWithConfig() echo.MiddlewareFunc {
  return middleware.CORSWithConfig(middleware.CORSConfig {
    AllowOrigins: []string{"http://localhost:8080", "https://tododot.site"},
    AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
    AllowCredentials: true,
  })
}
