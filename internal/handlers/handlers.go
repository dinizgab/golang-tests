package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dinizgab/golang-tests/internal/models"
	"github.com/dinizgab/golang-tests/internal/usecase"
)

func FindAllUsers(uc usecase.UserUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// code here
	}
}

func FindUserByID(uc usecase.UserUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// code here
	}
}

func CreateUser(uc usecase.UserUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := uc.Save(user); err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	}
}

func FollowUser(uc usecase.UserUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			FollowerID string `json:"follower_id"`
			FollowedID string `json:"followed_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := uc.FollowUser(req.FollowerID, req.FollowedID); err != nil {
			http.Error(w, "Failed to follow user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User followed successfully"})
	}
}
