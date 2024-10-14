package validators

import (
	"errors"
	"regexp"

	"github.com/Tes-sudo/online-learning-platform/user-service/models"
)


func ValidateUser(user *models.UserModel) error {
	if user.Username == "" {
		return errors.New("username is required")
	}
	if len(user.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}