package delivery_test

import (
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"

  "github.com/google/uuid"
  "github.com/gorilla/context"
  "github.com/gorilla/sessions"
  "github.com/labstack/echo/v4"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/require"
  "github.com/yinhr/todo-api/delivery"
  "github.com/yinhr/todo-api/middleware/session"
  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/usecase/signup"
  "github.com/yinhr/todo-api/util"
)

func TestSignupPost(t *testing.T) {
  u := model.User{
    ID: strings.Replace(uuid.New().String(), "-", "", -1),
    Email: "test@example.com",
    Password: "password001",
    PasswordConfirmation: "password001",
  }
  userJson, err := json.Marshal(u)
  assert.NoError(t, err)

  ucaseMock := signup.UsecaseMock{}
  ucaseMock.On("Store", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)

  e := echo.New()
  e.Use(session.SessionWithDefaultConfig(sessions.NewCookieStore([]byte("secret"))))

  req := util.NewHttpRequest(t, echo.POST, "/signup", strings.NewReader(string(userJson)))
  rec := httptest.NewRecorder()
  ctx := e.NewContext(req, rec)
  ctx.SetPath("/signup")
  ctx.Set("_session_store", sessions.NewCookieStore([]byte("secret")))
  defer context.Clear(ctx.Request())

  handler := &delivery.SignupHandler{&ucaseMock}
  err = handler.Post(ctx)
  require.NoError(t, err)

  cookies := rec.Result().Cookies()
  assert.True(t, contains(cookies, "sessionid"))
  //t.Log(cookies[0].Value)

  assert.Equal(t, http.StatusOK, rec.Code)
  ucaseMock.AssertExpectations(t)
}

func TestSignupBindError(t *testing.T) {
  e := echo.New()
  req := util.NewHttpRequest(t, echo.POST, "/signup", nil)
  rec := httptest.NewRecorder()
  ctx := e.NewContext(req, rec)
  ctx.SetPath("/signup")
  defer context.Clear(ctx.Request())

  handler := &delivery.SignupHandler{nil}
  err := handler.Post(ctx)
  assert.NoError(t, err)
  assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func contains(cookies []*http.Cookie, name string) bool {
  for _, v := range cookies {
    if v.Name == name {
      return true
    }
  }
  return false
}
