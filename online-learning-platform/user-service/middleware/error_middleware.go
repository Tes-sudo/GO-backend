package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Tes-sudo/online-learning-platform/user-service/errors"
	"github.com/Tes-sudo/online-learning-platform/user-service/logging"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logging.Error("Panic occurred: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

func RespondWithError(w http.ResponseWriter, err error, status int) {
	logging.Error("HTTP %d - %v", status, err)
	w.WriteHeader(status)
	response := ErrorResponse{Error: err.Error()}
	json.NewEncoder(w).Encode(response)
}

func HandleError(w http.ResponseWriter, err error) {
	switch err {
	// 400 Bad Request
	case errors.ErrInvalidInput,
		errors.ErrInvalidEmail,
		errors.ErrInvalidPassword,
		errors.ErrInvalidUsername,
		errors.ErrPasswordTooShort,
		errors.ErrUsernameTooShort,
		errors.ErrEmailRequired,
		errors.ErrPasswordRequired,
		errors.ErrUsernameRequired,
		errors.ErrInvalidDateFormat,
		errors.ErrInvalidPhoneNumber:
		RespondWithError(w, err, http.StatusBadRequest)

	// 401 Unauthorized
	case errors.ErrUnauthorized,
		errors.ErrInvalidCredentials,
		errors.ErrTokenExpired,
		errors.ErrInvalidToken:
		RespondWithError(w, err, http.StatusUnauthorized)

	// 404 Not Found
	case errors.ErrNotFound,
		errors.ErrUserNotFound:
		RespondWithError(w, err, http.StatusNotFound)

	// 409 Conflict
	case errors.ErrUserAlreadyExists,
		errors.ErrEmailAlreadyInUse,
		errors.ErrUsernameAlreadyInUse:
		RespondWithError(w, err, http.StatusConflict)

	// 429 Too Many Requests
	case errors.ErrRateLimitExceeded:
		RespondWithError(w, err, http.StatusTooManyRequests)

	// 500 Internal Server Error
	case errors.ErrInternal,
		errors.ErrDatabaseConnection,
		errors.ErrDatabaseQuery,
		errors.ErrTokenCreationFailed,
		errors.ErrRegistrationFailed,
		errors.ErrUpdateFailed,
		errors.ErrDeleteFailed:
		RespondWithError(w, errors.ErrInternal, http.StatusInternalServerError)

	// 503 Service Unavailable
	case errors.ErrServiceUnavailable:
		RespondWithError(w, err, http.StatusServiceUnavailable)

	// Default case for any unhandled errors
	default:
		RespondWithError(w, errors.ErrInternal, http.StatusInternalServerError)
	}
}
