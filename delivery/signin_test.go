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
  "github.com/yinhr/todo-api/usecase/signin"
)

func TestSigninPost(t *testing.T) {
  user := model.User{
    ID: strings.Replace(uuid.New().String(), "-", "", -1),
    Email: "test@example.com",
    Password: "password001",
  }
  userJson, err := json.Marshal(user)
  assert.NoError(t, err)

  ucaseMock := signin.UsecaseMock{}
  ucaseMock.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)

  e := echo.New()
  e.Use(session.SessionWithDefaultConfig(sessions.NewCookieStore([]byte("secret"))))
  req, err := http.NewRequest(echo.POST, "/signin", strings.NewReader(string(userJson)))
  assert.NoError(t, err)
  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

  rec := httptest.NewRecorder()
  ctx := e.NewContext(req, rec)
  ctx.SetPath("/signin")
  ctx.Set("_session_store", sessions.NewCookieStore([]byte("secret")))
  defer context.Clear(ctx.Request())

  handler := &delivery.SigninHandler{&ucaseMock}
  err = handler.Post(ctx)
  require.NoError(t, err)

  cookies := rec.Result().Cookies()
  assert.True(t, contains(cookies, "sessionid"))
  //t.Log(cookies[0].Value)

  assert.Equal(t, http.StatusOK, rec.Code)
  ucaseMock.AssertExpectations(t)
}
