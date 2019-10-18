package todo

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

func (m *UsecaseMock) Create(ctx context.Context, todo *model.Todo) error {
  ret := m.Called(ctx, todo)

  var r0 error
  if rf, ok := ret.Get(0).(func(context.Context, *model.Todo) error); ok {
    r0 = rf(ctx, todo)
  } else {
    r0 = ret.Error(0)
  }

  return r0
}
