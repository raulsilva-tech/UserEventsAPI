package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailIsRequired    = errors.New("email is required")
	ErrNameIsRequired     = errors.New("name is required")
	ErrPasswordIsRequired = errors.New("password is required")
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() error {

	if u.Name == "" {
		return ErrNameIsRequired
	}
	if u.Email == "" {
		return ErrEmailIsRequired
	}
	if u.Password == "" {
		return ErrPasswordIsRequired
	}

	return nil
}

func NewUser(name, email, password string, createdAt, updatedAt time.Time) (*User, error) {

	if password == "" {
		return nil, ErrPasswordIsRequired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Id:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  string(hash),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return u, u.Validate()
}
