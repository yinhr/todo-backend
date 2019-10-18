package model_test

import (
  "strings"
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/yinhr/todo-api/model"
  "gopkg.in/go-playground/validator.v9"
)

func TestUserModel(t *testing.T) {
  // valid user
  user := model.User{Email:"test@example.com", Password: strings.Repeat("a", 8), PasswordConfirmation: strings.Repeat("a", 8)}
  assert.Nil(t, validator.New().Struct(user))

  // valid user
  user = model.User{Email:"test@example.com", Password: strings.Repeat("a", 50), PasswordConfirmation: strings.Repeat("a", 50)}
  assert.Nil(t, validator.New().Struct(user))

  // invalid user (email blank)
  user = model.User{Email:"", Password: strings.Repeat("a", 9), PasswordConfirmation: strings.Repeat("a", 9)}
  assert.NotNil(t, validator.New().Struct(user))

  // invalid user (email with <>)
  user = model.User{Email:"te<st>@example.com", Password: strings.Repeat("a", 9), PasswordConfirmation: strings.Repeat("a", 9)}
  assert.NotNil(t, validator.New().Struct(user))

  // invalid user (passward blank)
  user = model.User{Email:"test@example.com", Password: "", PasswordConfirmation: ""}
  assert.NotNil(t, validator.New().Struct(user))

  // invalid user (passward not alphanum)
  user = model.User{Email:"test@example.com", Password: "abcde123<>", PasswordConfirmation: "abcde123<>"}
  assert.NotNil(t, validator.New().Struct(user))

  // invalid user (passward too long)
  user = model.User{Email:"test@example.com", Password: strings.Repeat("a", 51), PasswordConfirmation: strings.Repeat("a", 51)}
  assert.NotNil(t, validator.New().Struct(user))

  // invalid user (Passward neq PasswordConfirmation)
  user = model.User{Email:"test@example.com", Password: strings.Repeat("a", 10), PasswordConfirmation: strings.Repeat("b", 10)}
  assert.NotNil(t, validator.New().Struct(user))
}
