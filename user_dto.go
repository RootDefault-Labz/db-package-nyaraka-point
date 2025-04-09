package nyarakadb

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateUser struct {
	Username     string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Email        string `json:"email" validate:"required,email"`
	PasswordHash string `json:"password_hash" validate:"required,min=8"`
}

func (r *CreateUser) Validate() error {

	if err := validate.Struct(r); err != nil {
		return err
	}

	if !hasPasswordComplexity(r.PasswordHash) {
		return fmt.Errorf("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	return nil
}

type CheckUserExists struct {
	Username string `json:"username,omitempty" validate:"omitempty,min=3,max=50,alphanum"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
}

func (r *CheckUserExists) Validate() error {

	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.Username == "" && r.Email == "" {
		return fmt.Errorf("either username or email must be provided")
	}

	return nil
}

type UpdateUser struct {
	Username  string    `json:"username,omitempty" validate:"omitempty,min=3,max=50,alphanum"`
	Email     string    `json:"email,omitempty" validate:"omitempty,email"`
	IsActive  *bool     `json:"is_active,omitempty"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

func (r *UpdateUser) Validate() error {

	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.Username == "" && r.Email == "" && r.IsActive == nil {
		return fmt.Errorf("at least one field to update must be provided")
	}

	return nil
}

func hasPasswordComplexity(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' ||
			char == '^' || char == '&' || char == '*' || char == '(' || char == ')' ||
			char == '-':
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}