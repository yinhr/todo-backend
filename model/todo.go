package model

import (
  "time"
)

type (
  Todo struct {
    ID string `json:"id,omitempty"`
    UserID string `json:"user_id,omitempty"`
    Title string `json:"title" validate:"required"`
    Vital bool `json:"vital" validate:"required"`
    Done bool `json:"done" validate:"required"`
    Tasks []*Task `json:"tasks,omitempty"`
    Due time.Time `json:"due" validate:"required"`
    CreatedAt time.Time `json:"createdat" validate:"required"`
    UpdatedAt time.Time `json:"updatedat,omitempty"`
  }

  Task struct {
    ID string `json:"id,omitempty"`
    TodoID string `json:"todo_id,omitempty"`
    Title string `json:"title" validate:"required"`
    Done bool `json:"done" validate:"required"`
  }
)
