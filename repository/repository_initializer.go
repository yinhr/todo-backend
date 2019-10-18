package repository

import (
  "database/sql"

  "github.com/yinhr/todo-api/repository/user"
  "github.com/yinhr/todo-api/repository/todo"
  "github.com/yinhr/todo-api/util"
)

type (
  Repo int
  Repos map[Repo]interface{}
)

const (
  UserRepo Repo = iota
  TodoRepo
)

func GetInitializedRepositories(conn *sql.DB) Repos {
  repos := make(Repos)
  repos[UserRepo] = &user.UserRepository{conn, &util.Bcrypto{}, &util.UUID{}}
  repos[TodoRepo] = &todo.TodoRepository{conn, &util.UUID{}}
  return repos
}
