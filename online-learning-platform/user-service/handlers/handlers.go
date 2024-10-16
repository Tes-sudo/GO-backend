package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tes-sudo/online-learning-platform/user-service/auth"
	"github.com/Tes-sudo/online-learning-platform/user-service/errors"
	"github.com/Tes-sudo/online-learning-platform/user-service/middleware"
	"github.com/Tes-sudo/online-learning-platform/user-service/models"
	"github.com/Tes-sudo/online-learning-platform/user-service/repository"
	"github.com/Tes-sudo/online-learning-platform/user-service/validators"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	if err := validators.ValidateUser(&user); err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	if err := repository.CreateUser(&user); err != nil {
		middleware.HandleError(w, errors.ErrInternal)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	user, err := repository.GetUserByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			middleware.HandleError(w, errors.ErrNotFound)
		} else {
			middleware.HandleError(w, errors.ErrInternal)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	var updatedUser models.UserModel
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	if err := validators.ValidateUser(&updatedUser); err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	updatedUser.ID = uint(id)

	if err := repository.UpdateUser(&updatedUser); err != nil {
		if err == gorm.ErrRecordNotFound {
			middleware.HandleError(w, errors.ErrNotFound)
		} else {
			middleware.HandleError(w, errors.ErrInternal)
		}
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	if err := repository.DeleteUser(uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			middleware.HandleError(w, errors.ErrNotFound)
		} else {
			middleware.HandleError(w, errors.ErrInternal)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var loginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	// TODO: Validate credentials against database
	    user, err := repository.GetUserByEmail(loginRequest.Email)
	if err != nil || !validatePassword(user.Password, loginRequest.Password) {
		middleware.HandleError(w, errors.ErrUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		middleware.HandleError(w, errors.ErrInternal)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// TODO: Implement this function
func validatePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var registerRequest models.UserModel

	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		middleware.HandleError(w, errors.ErrInvalidInput)
		return
	}

	// Validate user input
	if err := validators.ValidateUser(&registerRequest); err != nil {
		middleware.HandleError(w, err)
		return
	}

	// Check if user already exists
	existingUser, err := repository.GetUserByEmail(registerRequest.Email)
	if err == nil && existingUser != nil {
		middleware.HandleError(w, errors.ErrUserAlreadyExists)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		middleware.HandleError(w, errors.ErrInternal)
		return
	}
	registerRequest.Password = string(hashedPassword)

	// Create user
	if err := repository.CreateUser(&registerRequest); err != nil {
		middleware.HandleError(w, errors.ErrInternal)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
