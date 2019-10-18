package user

import (
  "context"
  "database/sql"
  "strings"

  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/util"
  log "github.com/sirupsen/logrus"
)

type (
  UserRepository struct {
    Conn *sql.DB
    Bcrypt util.Crypto
    UUID util.UID
  }
)

func (u *UserRepository) FindBy(ctx context.Context, email string, password string) (*model.User, error) {
  queryStr := `SELECT LOWER(HEX(id)), password_digest FROM users WHERE email=? LIMIT 1`
  query, err := u.Conn.PrepareContext(ctx, queryStr)
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  rows, err := query.QueryContext(ctx, strings.ToLower(email))
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  defer rows.Close()
  user := new(model.User)
  if rows.Next() {
    err = rows.Scan(&user.ID, &user.Password)
    if err != nil {
      log.Error(err)
      return nil, util.NewInternalServerError()
    }
  } else {
    return nil, util.NewBadRequestError("メールアドレス、またはパスワードが間違っています")
  }
  err = u.Bcrypt.CompareHashAndPassword(user.Password, password)
  if err != nil {
    log.Error(err)
    return nil, util.NewBadRequestError("メールアドレス、またはパスワードが間違っています")
  }
  user.Password = ""
  return user, nil
}

func (u *UserRepository) Fetch(ctx context.Context) ([]*model.User, error) {
  query := `SELECT LOWER(HEX(id)), email FROM users`
  users := make([]*model.User, 0)
  rows, err := u.Conn.QueryContext(ctx, query)
  if err != nil {
    log.Error(err)
    return nil, util.NewInternalServerError()
  }
  for rows.Next() {
    u := new(model.User)
    err = rows.Scan(&u.ID, &u.Email)
    if err != nil {
      log.Error(err)
      return nil, util.NewInternalServerError()
    }
    users = append(users, u)
  }
  return users, nil
}

func (u *UserRepository) Store(ctx context.Context, user *model.User) error {
  queryStr := `INSERT INTO users(id, email, password_digest) VALUES(UNHEX(REPLACE(?, "-", "")), ?, ?)`
  query, err := u.Conn.PrepareContext(ctx, queryStr)
  if err != nil {
    log.Error(err)
    return util.NewInternalServerError()
  }
  id := u.UUID.New()
  hash, err := u.Bcrypt.GenerateFromPassword(user.Password)
  _, err = query.ExecContext(ctx, id, strings.ToLower(user.Email), string(hash))
  if err != nil {
    log.Error(err)
    return util.SqlErrorHandler(err)
  }
  user.ID = strings.Replace(u.UUID.String(id), "-", "", -1)
  return nil
}
