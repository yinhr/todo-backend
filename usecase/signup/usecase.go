package signup

import (
  "context"

  "github.com/yinhr/todo-api/model"
)

type (
  Usecase interface {
    Fetch(ctx context.Context) ([]*model.User, error)
    Store(ctx context.Context, user *model.User) error
  }
)
