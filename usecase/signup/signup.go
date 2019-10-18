package signup

import (
  "context"

  "github.com/yinhr/todo-api/repository/user"
  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/util"
  log "github.com/sirupsen/logrus"
  "gopkg.in/go-playground/validator.v9"
)

type (
  SignupUsecase struct {
    UserRepository user.Repository
  }
)

func (u *SignupUsecase) Fetch(ctx context.Context) ([]*model.User, error) {
  users, _ := u.UserRepository.Fetch(ctx)
  return users, nil
}

func (u *SignupUsecase) Store(ctx context.Context, user *model.User) error {
  if err := validator.New().Struct(user); err != nil {
    log.Error(err)
    return util.NewBadRequestError("400: Bad Request")
  }
  if err := u.UserRepository.Store(ctx, user); err != nil {
    return err
  }
  return nil
}
