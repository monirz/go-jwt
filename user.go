package gojwt

import "time"

type User struct {
	ID        int64     `json:"-"`
	UUID      string    `json:"uuid"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserService interface {
	CreateUser(user *User) (int64, error)
	FindUserByID(userID string) (*User, error)
	FindByEmail(email string) (*User, error)
	UpdateUser(user *User) error
}
