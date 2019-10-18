package util

import (
  "os"
)

func Production() bool {
  return "production" == os.Getenv("ECHO_ENV")
}

func Development() bool {
  return "develepment" == os.Getenv("ECHO_ENV")
}
