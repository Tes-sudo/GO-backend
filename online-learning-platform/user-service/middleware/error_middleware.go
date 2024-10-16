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
	case errors.ErrInvalidInput:
		RespondWithError(w, err, http.StatusBadRequest)
	case errors.ErrNotFound:
		RespondWithError(w, err, http.StatusNotFound)
	case errors.ErrUnauthorized:
		RespondWithError(w, err, http.StatusUnauthorized)
	default:
		RespondWithError(w, errors.ErrInternal, http.StatusInternalServerError)
	}
}
