package signin

import (
  "context"
  "github.com/stretchr/testify/mock"
  "github.com/yinhr/todo-api/model"
)

type (
  UsecaseMock struct {
    mock.Mock
  }
)

func (m *UsecaseMock) Create(ctx context.Context, user *model.User) error {
  ret := m.Called(ctx, user)

  var r0 error
  if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
    r0 = rf(ctx, user)
  } else {
    r0 = ret.Error(0)
  }

  return r0
}
