package User

type User struct {
	Username string `json:"Username" validate:"required,alphanum"`
	Email    string `json:"Email" validate:"required,email"`
}
