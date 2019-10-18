package util

import (
  "database/sql"
  "testing"
  "github.com/DATA-DOG/go-sqlmock"
)

func NewSQLMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
  t.Helper()
  db, mock, err := sqlmock.New()
  if err != nil {
    t.Fatalf("expected no error, but got:'%s'", err)
  }
  return db, mock
}
