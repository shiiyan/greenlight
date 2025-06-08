package data

import (
	"errors"
	"time"

	"github.com/shiiyan/greenlight/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64
	CreatedAt time.Time
	Name      string
	Email     string
	Password  password
	Activated bool
	Version   int
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(user.Email != "", "email", "must be provided")

	v.Check(*user.Password.plaintext != "", "password", "must be provided")
	v.Check(len(*user.Password.plaintext) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(*user.Password.plaintext) <= 72, "password", "must not be more than 72 bytes long")

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}
