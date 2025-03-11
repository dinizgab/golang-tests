package handlers

import (
	"net/http"

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
		// code here
	}
}

func FollowUser(uc usecase.UserUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// code here
	}
}
