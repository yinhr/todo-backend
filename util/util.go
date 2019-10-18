package util

import (
  "github.com/google/uuid"
)

type (
  Crypto interface {
    GenerateFromPassword(password string) (string, error)
    CompareHashAndPassword(hashedPassword, password string) error
  }
  UID interface {
    New() uuid.UUID
    Parse(src string) (uuid.UUID, error)
    String(src uuid.UUID) string
  }
)

type (
  QueryParams struct {
    UserID string `json:"userID" query:"userID"`
    Cursor int `json:"cursor" query:"cursor"`
    OrderBy string `json:"orderBy" query:"orderBy"`
    Direction string `json:"direction" query:"direction"`
  }
)
