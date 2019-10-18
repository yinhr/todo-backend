package todo

import (
  "context"

  "github.com/yinhr/todo-api/repository/todo"
  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/util"
  log "github.com/sirupsen/logrus"
)

type (
  TodoUsecase struct {
    TodoRepository todo.Repository
  }
)

func (u *TodoUsecase) FindBy(ctx context.Context, id string) (*model.Todo, error) {
  todo, err := u.TodoRepository.FindBy(ctx, id)
  if err != nil {
    return nil, err
  }
  return todo, nil
}

func (u *TodoUsecase) Fetch(ctx context.Context, params *util.QueryParams) ([]*model.Todo, error) {
  if params.OrderBy != "due" && params.OrderBy != "created_at" {
    log.Error("Bad request query 'orderBy'")
    return nil, util.NewBadRequestError("400: Bad Request")
  }
  if params.Direction != "desc" && params.Direction != "asc" {
    log.Error("Bad request query 'direction'")
    return nil ,util.NewBadRequestError("400: Bad Request")
  }
  //todos, err := u.TodoRepository.Fetch(ctx, params.Cursor, params.OrderBy, params.Direction)
  todos, err := u.TodoRepository.Fetch(ctx, params)
  if err != nil {
    return nil, util.NewBadRequestError("400: Bad Request")
  }
  return todos, nil
}

func (u *TodoUsecase) Create(ctx context.Context, todo *model.Todo) error {
  if err := u.TodoRepository.Store(ctx, todo); err != nil {
    return err
  }
  return nil
}

func (u *TodoUsecase) PatchTodoDone(ctx context.Context, todo *model.Todo) error {
  if err := u.TodoRepository.PatchTodoDone(ctx, todo); err != nil {
    return err
  }
  return nil
}

func (u *TodoUsecase) PatchTaskDone(ctx context.Context, task *model.Task) error {
  if err := u.TodoRepository.PatchTaskDone(ctx, task); err != nil {
    return err
  }
  return nil
}

func (u *TodoUsecase) Put(ctx context.Context, todo *model.Todo) error {
  if err := u.TodoRepository.Put(ctx, todo); err != nil {
    return err
  }
  return nil
}

func (u *TodoUsecase) Delete(ctx context.Context, id string) error {
  if err := u.TodoRepository.Delete(ctx, id); err != nil {
    return err
  }
  return nil
}
