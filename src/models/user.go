package models

import (
	"errors"
	"strings"

	"github.com/margen2/shorgot/src/security"

	"github.com/badoux/checkmail"
)

// User represents a user on the database
type User struct {
	ID       uint64 `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// Prepare verifies if the user information is valid for the current stage
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}
	if err := user.format(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.Email == "" {
		return errors.New("email can't be null value")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return err
	}
	if stage == "signup" && user.Password == "" {
		return errors.New("password can't be null value")
	}
	return nil
}

func (user *User) format(stage string) error {
	user.Email = strings.TrimSpace(user.Email)

	if stage == "signup" {
		hashedPW, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPW)
	}
	return nil
}
