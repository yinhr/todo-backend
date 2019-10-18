package user

import (
  "context"

  "github.com/yinhr/todo-api/model"
)

type (
  Repository interface {
    FindBy(ctx context.Context, email string, password string) (*model.User, error)
    Fetch(ctx context.Context) ([]*model.User, error)
    Store(ctx context.Context, user *model.User) error
  }
)
