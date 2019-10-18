package user

import (
  "context"
  "strings"
  "testing"
  "regexp"

  "github.com/DATA-DOG/go-sqlmock"
  "github.com/stretchr/testify/assert"
  "github.com/yinhr/todo-api/model"
  "github.com/yinhr/todo-api/util"
)

/*
func TestFetch(t *testing.T) {
  db, mock, err := sqlmock.New()
  if err != nil {
    t.Fatalf("Error: '%s' was not expected when opening mock", err)
  }
  defer db.Close()
  mockUsers := []model.User {
    model.User {Email: "test@example.com"},
  }
  rows := NewRows([]string{"email"}).Addrow(mockUsers[0].Email)
  query := `SELECT email FROM users`
  mock.ExpectQuery(query).WillReturnRows(rows)
  mock.ExpectBegin()
}
*/

func TestFindBy(t *testing.T) {
  email := "test@example.com"
  passrowd := "password"
  password_digest := "password_digest"
  db, mock := util.NewSQLMock(t)
  defer db.Close()
  userRepo := &UserRepository{db, &util.CryptoMock{}, &util.UUIDMock{}}
  uid := userRepo.UUID.New()
  id := strings.Replace(userRepo.UUID.String(uid), "-", "", -1)

  mock.ExpectPrepare(regexp.QuoteMeta(`SELECT LOWER(HEX(id)), password_digest FROM users WHERE email=? LIMIT 1`))
  mock.ExpectQuery(regexp.QuoteMeta(`SELECT LOWER(HEX(id)), password_digest FROM users WHERE email=? LIMIT 1`)).
    WithArgs(email).
    WillReturnRows(sqlmock.NewRows([]string{"id", "password_digest"}).AddRow(id, password_digest))
  user, err := userRepo.FindBy(context.Background(), email, passrowd)
  t.Log(err)
  assert.NoError(t, err)
  assert.NotNil(t, user)
  assert.Equal(t, "", user.Password)
  assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStore(t *testing.T) {
  db, mock := util.NewSQLMock(t)
  defer db.Close()
  userRepo := &UserRepository{db, &util.CryptoMock{}, &util.UUIDMock{}}
  user := &model.User{
    Email: "test@example.com", 
    Password: "passrowd",
  }
  id := userRepo.UUID.New()

  mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO users(id, email, password_digest) VALUES(UNHEX(REPLACE(?, "-", "")), ?, ?)`))
  mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users(id, email, password_digest) VALUES(UNHEX(REPLACE(?, "-", "")), ?, ?)`)).
    WithArgs(id, "test@example.com", "password_digest").
    WillReturnResult(sqlmock.NewResult(1, 1)) // NewResult('LastInsertId', 'NumberOfRowsAffected')
  assert.NoError(t, userRepo.Store(context.Background(), user))
  assert.NoError(t, mock.ExpectationsWereMet())
}

