package todo

import (
  "context"

  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/util"
)

type (
  Repository interface {
    FindBy(ctx context.Context, id string) (*model.Todo, error)
    Fetch(ctx context.Context, params *util.QueryParams) ([]*model.Todo, error)
    Store(ctx context.Context, todo *model.Todo) error
    PatchTodoDone(ctx context.Context, todo *model.Todo) error
    PatchTaskDone(ctx context.Context, task *model.Task) error
    Put(ctx context.Context, todo *model.Todo) error
    Delete(ctx context.Context, id string) error
  }
)
