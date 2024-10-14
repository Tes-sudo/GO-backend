package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tes-sudo/online-learning-platform/user-service/models"
	"github.com/Tes-sudo/online-learning-platform/user-service/repository"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
		}
	var updatedUser models.UserModel
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		}
		updatedUser.ID = uint(id)
		err = repository.UpdateUser(&updatedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUserHandler(w http.ResponseWriter,r *http.Request){
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil{
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	err = repository.DeleteUser(uint(id))
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
		
}




// TODO: Add UpdateUserHandler and DeleteUserHandler