package signin

import (
  "context"

  "github.com/yinhr/todo-api/repository/user"
  "github.com/yinhr/todo-api/model"
  //log "github.com/sirupsen/logrus"
)

type (
  SigninUsecase struct {
    UserRepository user.Repository
  }
)

func (u *SigninUsecase) Create(ctx context.Context, user *model.User) error {
  if u, err := u.UserRepository.FindBy(ctx, user.Email, user.Password); err != nil {
    return err
  } else {
    user.ID = u.ID
    return nil
  }
}

