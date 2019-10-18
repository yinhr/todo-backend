package signin

import (
  "context"

  "github.com/yinhr/todo-api/model"
)

type (
  Usecase interface {
    Create(ctx context.Context, user *model.User) error
  }
)
