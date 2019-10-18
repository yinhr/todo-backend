package signup_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/repository/user"
  "github.com/yinhr/todo-api/usecase/signup"
  "github.com/yinhr/todo-api/util"
)

func TestSignupStore(t *testing.T) {
  u := model.User{
    Email: "test@example.com",
    Password: "password1",
    PasswordConfirmation: "password1",
  }
  userMock := user.RepositoryMock{}
  userMock.On("Store", mock.Anything, &u).Return(nil)

  signupUsecase := &signup.SignupUsecase{&userMock}
  assert.NoError(t, signupUsecase.Store(nil, &u))
}

func TestSignupStoreError(t *testing.T) {
  u := model.User{
    Email: "test@example.com",
    Password: "password1",
    PasswordConfirmation: "password1",
  }
  userMock := user.RepositoryMock{}
  userMock.On("Store", mock.Anything, &u).Return(util.NewInternalServerError())

  signupUsecase := &signup.SignupUsecase{&userMock}
  assert.Error(t, signupUsecase.Store(nil, &u))
}

func TestSignupValidateError(t *testing.T) {
  u := model.User{
    Email: "test@example.com",
    Password: "aaa",
    PasswordConfirmation: "bbb",
  }
  signupUsecase := &signup.SignupUsecase{nil}
  assert.Error(t, signupUsecase.Store(nil, &u))
}
