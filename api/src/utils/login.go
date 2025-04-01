package utils

import (
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginBody) Validate() error {
	if req.Email == "" {
		return errors.New("email is required")
	}

	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}
