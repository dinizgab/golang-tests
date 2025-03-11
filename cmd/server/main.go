package main

import (
	"log"
	"net/http"

	"github.com/dinizgab/golang-tests/internal/config"
	"github.com/dinizgab/golang-tests/internal/db"
	"github.com/dinizgab/golang-tests/internal/handlers"
	"github.com/dinizgab/golang-tests/internal/repository"
	"github.com/dinizgab/golang-tests/internal/service"
	"github.com/dinizgab/golang-tests/internal/usecase"
)

func main() {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	broker, err := service.NewBrokerConnection()
	if err != nil {
		log.Fatal(err)
	}
	notificationService := service.NewNotificationService(broker)

	userRepository := repository.NewUserRepository(db)
    userUsecase := usecase.NewUserUsecase(userRepository, notificationService)

    http.HandleFunc("GET /users", handlers.FindAllUsers(userUsecase))
    http.HandleFunc("GET /users/{id}", handlers.FindUserByID(userUsecase))
    http.HandleFunc("POST /users", handlers.CreateUser(userUsecase))
    http.HandleFunc("POST /users/follow", handlers.FollowUser(userUsecase))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
