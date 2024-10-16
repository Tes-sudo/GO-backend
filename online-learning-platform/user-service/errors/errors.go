package errors

import "errors"

var (
	// General errors
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("resource not found")
	ErrInternal     = errors.New("internal server error")
	ErrUnauthorized = errors.New("unauthorized")

	// User-specific errors
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidEmail        = errors.New("invalid email format")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInvalidUsername     = errors.New("invalid username")
	ErrPasswordTooShort    = errors.New("password is too short")
	ErrUsernameTooShort    = errors.New("username is too short")
	ErrEmailRequired       = errors.New("email is required")
	ErrPasswordRequired    = errors.New("password is required")
	ErrUsernameRequired    = errors.New("username is required")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrEmailAlreadyInUse   = errors.New("email is already in use")
	ErrUsernameAlreadyInUse = errors.New("username is already in use")

	// Authentication errors
	ErrTokenExpired        = errors.New("token has expired")
	ErrInvalidToken        = errors.New("invalid token")
	ErrTokenCreationFailed = errors.New("failed to create token")

	// Database errors
	ErrDatabaseConnection = errors.New("database connection error")
	ErrDatabaseQuery      = errors.New("database query error")

	// Validation errors
	ErrInvalidDateFormat  = errors.New("invalid date format")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")

	// Registration errors
	ErrRegistrationFailed = errors.New("registration failed")

	// Update errors
	ErrUpdateFailed = errors.New("update failed")

	// Delete errors
	ErrDeleteFailed = errors.New("delete failed")

	// Other potential errors
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
	ErrServiceUnavailable = errors.New("service temporarily unavailable")
)
