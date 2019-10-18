package user

import (
  "context"
  "github.com/stretchr/testify/mock"
  "github.com/yinhr/todo-api/model"
)

type (
  RepositoryMock struct {
    mock.Mock
  }
)

func (m *RepositoryMock) FindBy(ctx context.Context, email string, password string) (*model.User, error) {
  ret := m.Called(ctx, email, password)

  var r0 *model.User
  if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.User); ok {
    r0 = rf(ctx, email, password)
  } else {
    if ret.Get(0) != nil {
      r0 = ret.Get(0).(*model.User)
    }
  }

  var r1 error
  if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
    r1 = rf(ctx, email, password)
  } else {
    r1 = ret.Error(1)
  }

  return r0, r1
}

func (m *RepositoryMock) Fetch(ctx context.Context) ([]*model.User, error) {
  ret := m.Called(ctx)

  var r0 []*model.User
  if rf, ok := ret.Get(0).(func(context.Context) []*model.User); ok {
    r0 = rf(ctx)
  } else {
    if ret.Get(0) != nil {
      r0 = ret.Get(0).([]*model.User)
    }
  }

  var r1 error
  if rf, ok := ret.Get(1).(func(context.Context) error); ok {
    r1 = rf(ctx)
  } else {
    r1 = ret.Error(1)
  }

  return r0, r1
}

func (m *RepositoryMock) Store(ctx context.Context, user *model.User) error {
  ret := m.Called(ctx, user)

  var r0 error
  if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
    r0 = rf(ctx, user)
  } else {
    r0 = ret.Error(0)
  }

  return r0
}
