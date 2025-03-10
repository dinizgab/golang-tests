package main

import (
	"log"

	"github.com/dinizgab/golang-tests/internal/config"
	"github.com/dinizgab/golang-tests/internal/db"
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

}
