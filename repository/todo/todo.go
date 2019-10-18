package todo

import (
  "context"
  "database/sql"

  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/util"
  log "github.com/sirupsen/logrus"
)

type (
  TodoRepository struct {
    Conn *sql.DB
    UUID util.UID
  }
)

func (t *TodoRepository) FindBy(ctx context.Context, id string) (*model.Todo, error) {
  queryStr := `SELECT LOWER(HEX(id)), title, vital, due FROM todos WHERE id = UNHEX(?)`
  query, err := t.Conn.PrepareContext(ctx, queryStr)
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  rows, err := query.QueryContext(ctx, id)
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  defer rows.Close()
  todo := new(model.Todo)
  if rows.Next() {
    err = rows.Scan(&todo.ID, &todo.Title, &todo.Vital, &todo.Due)
    if err != nil {
      log.Error(err)
      return nil, util.NewInternalServerError()
    }
    todo.Tasks, err = t.getTasks(ctx, todo.ID)
    if err != nil {
      log.Error(err)
      return nil, util.NewInternalServerError()
    }
  }
  return todo, nil
}

func (t *TodoRepository) Fetch(ctx context.Context, params *util.QueryParams) ([]*model.Todo, error) {
  queryStr := `SELECT LOWER(HEX(id)), title, vital, done, due, created_at FROM todos ` +
              `WHERE user_id = UNHEX(?) ORDER BY ` + params.OrderBy + ` ` + params.Direction + `, id DESC LIMIT 5 OFFSET ?`
  query, err := t.Conn.PrepareContext(ctx, queryStr)
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  rows, err := query.QueryContext(ctx, params.UserID, params.Cursor)
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  defer rows.Close()
  todos := make([]*model.Todo, 0)
  for rows.Next() {
    todo := new(model.Todo)
    err = rows.Scan(&todo.ID, &todo.Title, &todo.Vital, &todo.Done, &todo.Due, &todo.CreatedAt)
    if err != nil {
      log.Error(err)
      return nil, util.NewInternalServerError()
    }
    todo.Tasks, err = t.getTasks(ctx, todo.ID)
    if err != nil {
      log.Error(err)
      return nil, err
    }
    todos = append(todos, todo)
  }
  return todos, nil
}

func (t *TodoRepository) Store(ctx context.Context, todo *model.Todo) error {
  tx, err := t.Conn.Begin()
  if err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  trans := func(tx *sql.Tx, todo *model.Todo) error {
    queryStr := `INSERT INTO todos(id, user_id, title, vital, done, due) VALUES(UNHEX(REPLACE(?, "-", "")), UNHEX(REPLACE(?, "-", "")), ?, ?, ?, ?)`
    query, err := tx.PrepareContext(ctx, queryStr)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    id := t.UUID.New()
    uid, err := t.UUID.Parse(todo.UserID)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    _, err = query.ExecContext(ctx, id, uid, todo.Title, todo.Vital, todo.Done, todo.Due)
    if err != nil {
      log.Error(err)
      return util.SqlErrorHandler(err)
    }
    queryStr = `INSERT INTO tasks(id, todo_id, title, done) VALUES(UNHEX(REPLACE(?, "-", "")), UNHEX(REPLACE(?, "-", "")), ?, ?)`
    query, err = tx.PrepareContext(ctx, queryStr)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    for _, task := range todo.Tasks {
      _, err := query.ExecContext(ctx, t.UUID.New(), id, task.Title, task.Done)
      if err != nil {
        log.Error(err)
        return util.SqlErrorHandler(err)
      }
    }
    return nil
  }
  if err := trans(tx, todo); err != nil {
    if re := tx.Rollback(); re != nil {
      log.Error(re)
      return util.NewInternalServerError()
    }
    log.Error(err)
    return util.NewInternalServerError()
  }
  if err := tx.Commit(); err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  return nil
}

func (t *TodoRepository) PatchTaskDone(ctx context.Context, task *model.Task) error {
  tx, err := t.Conn.Begin()
  if err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  trans := func(tx *sql.Tx, task *model.Task) error {
    queryStr := `UPDATE tasks SET done=? WHERE id = UNHEX(?)` 
    query, err := tx.PrepareContext(ctx, queryStr)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    _, err = query.ExecContext(ctx, task.Done, task.ID)
    if err != nil {
      log.Error(err)
      return util.SqlErrorHandler(err)
    }
    return nil
  }
  if err := trans(tx, task); err != nil {
    if re := tx.Rollback(); re != nil {
      log.Error(re)
      return util.NewInternalServerError()
    }
    log.Error(err)
    return util.NewInternalServerError()
  }
  if err := tx.Commit(); err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  return nil
}

func (t *TodoRepository) PatchTodoDone(ctx context.Context, todo *model.Todo) error {
  tx, err := t.Conn.Begin()
  if err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  trans := func(tx *sql.Tx, todo *model.Todo) error {
    queryStr := `UPDATE todos SET done=? WHERE id = UNHEX(?)` 
    query, err := tx.PrepareContext(ctx, queryStr)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    _, err = query.ExecContext(ctx, todo.Done, todo.ID)
    if err != nil {
      log.Error(err)
      return util.SqlErrorHandler(err)
    }
    return nil
  }
  if err := trans(tx, todo); err != nil {
    if re := tx.Rollback(); re != nil {
      log.Error(re)
      return util.NewInternalServerError()
    }
    log.Error(err)
    return util.NewInternalServerError()
  }
  if err := tx.Commit(); err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  return nil
}

func (t *TodoRepository) Put(ctx context.Context, todo *model.Todo) error {
  tx, err := t.Conn.Begin()
  if err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  trans := func(tx *sql.Tx, todo *model.Todo) error {
    queryStr := `UPDATE todos SET title=?, due=?, done=?, vital=? WHERE id = UNHEX(?)` 
    query, err := tx.PrepareContext(ctx, queryStr)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    _, err = query.ExecContext(ctx, todo.Title, todo.Due, todo.Done, todo.Vital, todo.ID)
    if err != nil {
      log.Error(err)
      return util.SqlErrorHandler(err)
    }
    queryStr = `DELETE FROM tasks WHERE todo_id = UNHEX(?)` 
    query, err = tx.PrepareContext(ctx, queryStr)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    _, err = query.ExecContext(ctx, todo.ID)
    if err != nil {
      log.Error(err)
      return util.SqlErrorHandler(err)
    }
    queryStr = `INSERT INTO tasks(id, todo_id, title, done) VALUES(UNHEX(REPLACE(?, "-", "")), UNHEX(?), ?, ?)`
    query, err = tx.PrepareContext(ctx, queryStr)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    for _, task := range todo.Tasks {
      _, err = query.ExecContext(ctx, t.UUID.New(), todo.ID, task.Title, task.Done)
      if err != nil {
        log.Error(err)
        return util.SqlErrorHandler(err)
      }
    }
    return nil
  }
  if err := trans(tx, todo); err != nil {
    if re := tx.Rollback(); re != nil {
      log.Error(re)
      return util.NewInternalServerError()
    }
    log.Error(err)
    return util.NewInternalServerError()
  }
  if err := tx.Commit(); err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  return nil
}

func (t *TodoRepository) Delete(ctx context.Context, id string) error {
  tx, err := t.Conn.Begin()
  if err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  trans := func(tx *sql.Tx, id string) error {
    queryStr := `DELETE FROM todos WHERE id = UNHEX(?)` 
    query, err := tx.PrepareContext(ctx, queryStr)
    if err != nil {
      log.Error(err)
      return util.NewInternalServerError()
    }
    _, err = query.ExecContext(ctx, id)
    if err != nil {
      log.Error(err)
      return util.SqlErrorHandler(err)
    }
    return nil
  }
  if err := trans(tx, id); err != nil {
    if re := tx.Rollback(); re != nil {
      log.Error(re)
      return util.NewInternalServerError()
    }
    log.Error(err)
    return util.NewInternalServerError()
  }
  if err := tx.Commit(); err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  return nil
}

func (t *TodoRepository) getTasks(ctx context.Context, id string) ([]*model.Task, error) {
  queryStr := `SELECT LOWER(HEX(id)), title, done FROM tasks WHERE todo_id = UNHEX(?)`
  query, err := t.Conn.PrepareContext(ctx, queryStr)
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  rows, err := query.QueryContext(ctx, id)
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  defer rows.Close()
  tasks := make([]*model.Task, 0)
  for rows.Next() {
    task := new(model.Task)
    err := rows.Scan(&task.ID, &task.Title, &task.Done)
    if err != nil {
      log.Error(err)
      return nil, util.NewInternalServerError()
    }
    tasks = append(tasks, task)
  }
  return tasks, nil
}
