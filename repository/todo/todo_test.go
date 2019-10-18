package todo

import (
  "context"
  "testing"
  "time"
  "strings"
  "regexp"

  "github.com/DATA-DOG/go-sqlmock"
  "github.com/stretchr/testify/assert"
  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/util"
)

func TestStore(t *testing.T) {
  db, mock := util.NewSQLMock(t)
  defer db.Close()
  todoRepo := &TodoRepository{db, &util.UUIDMock{}}

  id := todoRepo.UUID.New()
  userID := todoRepo.UUID.New()
  title := "todo title"
  vital := true
  done := false
  due := time.Now()
  todo := &model.Todo{
    UserID: strings.Replace(todoRepo.UUID.String(userID), "-", "", -1),
    Title: title,
    Vital: vital,
    Done: done,
    Due: due,
  }

  mock.ExpectBegin()
  mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO todos(id, user_id, title, vital, done, due) VALUES(UNHEX(REPLACE(?, "-", "")), UNHEX(REPLACE(?, "-", "")), ?, ?, ?, ?)`))
  mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO todos(id, user_id, title, vital, done, due) VALUES(UNHEX(REPLACE(?, "-", "")), UNHEX(REPLACE(?, "-", "")), ?, ?, ?, ?)`)).
    WithArgs(id, userID, todo.Title, todo.Vital, todo.Done, todo.Due).
    WillReturnResult(sqlmock.NewResult(1, 1)) // NewResult('LastInsertId', 'NumberOfRowsAffected')

  tasks := []*model.Task {
    &model.Task{Title: "Task1", Done: false},
    &model.Task{Title: "Task2", Done: false},
  }
  todo.Tasks = tasks
  mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO tasks(id, todo_id, title, done) VALUES(UNHEX(REPLACE(?, "-", "")), UNHEX(REPLACE(?, "-", "")), ?, ?)`))
  for _, task := range tasks {
    mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO tasks(id, todo_id, title, done) VALUES(UNHEX(REPLACE(?, "-", "")), UNHEX(REPLACE(?, "-", "")), ?, ?)`)).
      WithArgs(todoRepo.UUID.New(), id, task.Title, task.Done).
      WillReturnResult(sqlmock.NewResult(1, 1)) // NewResult('LastInsertId', 'NumberOfRowsAffected')
  }

  mock.ExpectCommit()
  assert.NoError(t, todoRepo.Store(context.Background(), todo))
  assert.NoError(t, mock.ExpectationsWereMet())
}

