package usecase

import (
  "github.com/yinhr/todo-api/repository"
  "github.com/yinhr/todo-api/repository/user"
  rtodo "github.com/yinhr/todo-api/repository/todo"
  "github.com/yinhr/todo-api/usecase/signup"
  "github.com/yinhr/todo-api/usecase/signin"
  utodo "github.com/yinhr/todo-api/usecase/todo"
)

type (
  UCase int
  UCases map[UCase]interface{}
)

const (
  SignupUCase UCase = iota
  SigninUCase
  TodoUCase
)

func GetInitializedUsecases(repos *repository.Repos) UCases {
  ucases := make(UCases)
  userRepository, _ := (*repos)[repository.UserRepo].(*user.UserRepository)
  todoRepository, _ := (*repos)[repository.TodoRepo].(*rtodo.TodoRepository)
  ucases[SignupUCase] = &signup.SignupUsecase{userRepository}
  ucases[SigninUCase] = &signin.SigninUsecase{userRepository}
  ucases[TodoUCase] = &utodo.TodoUsecase{todoRepository}
  return ucases
}

