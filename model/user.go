package model

type (
  User struct {
    ID string `json:"id,omitempty"`
    //ID int64 `json:"id,omitempty"`
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,alphanumunicode,min=8,max=50,eqfield=PasswordConfirmation"`
    PasswordConfirmation string `json:"password_confirmation"`
    ImagePath string `json:"image_path,omitempty" validate:"omitempty,url"`
  }
)
