package util

import (
  "golang.org/x/crypto/bcrypt"
)

type (
  Bcrypto struct{}
  CryptoMock struct{}
)

func (b *Bcrypto) GenerateFromPassword(password string) (string, error) {
  if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
    return "", err
  } else {
    return string(hash), nil
  }
}

func (b *Bcrypto) CompareHashAndPassword(hashedPassword, password string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (m *CryptoMock) GenerateFromPassword(password string) (string, error) {
  return "password_digest", nil
}

func (m *CryptoMock) CompareHashAndPassword(hashedPassword, password string) error {
  return nil
}
