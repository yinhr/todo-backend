package util

import (
  "github.com/google/uuid"
)

var (
  uid = uuid.New()
  uidStr = uid.String()
)

type (
  UUID struct{}
  UUIDMock struct {}
)

func (u *UUID) New() uuid.UUID {
  return uuid.New()
}

func (u *UUID) Parse(src string) (uuid.UUID, error) {
  return uuid.Parse(src)
}

func (u *UUID) String(src uuid.UUID) string {
  return src.String()
}

func (m *UUIDMock) New() uuid.UUID {
  return uid
}

func (m *UUIDMock) Parse(src string) (uuid.UUID, error) {
  return uid, nil
}

func (m *UUIDMock) String(src uuid.UUID) string {
  return uidStr
}

