package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dinizgab/golang-tests/internal/db"
	"github.com/dinizgab/golang-tests/internal/models"
)

func TestRepository(t *testing.T) {
	dbConfig, err := db.NewDBConfig()
	if err != nil {
		t.Fatal(err)
	}
	db, err := db.New(dbConfig)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepository(db)

	t.Run("Test create new user", func(t *testing.T) {
		db.Exec("DELETE FROM users")

		user := models.User{
			FirstName: "Gabriel",
			Username:  "dinizgab",
		}

		err = repo.Save(user)

		assert.NoError(t, err)

        users, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Equal(t, "Gabriel", users[0].FirstName)
		assert.Equal(t, "dinizgab", users[0].Username)

	})

	t.Run("Test find all users", func(t *testing.T) {
		db.Exec("DELETE FROM users")

		repo.Save(
			models.User{
				FirstName: "Gabriel",
				Username:  "dinizgab",
			},
		)
		repo.Save(
			models.User{
				FirstName: "Joao",
				Username:  "john",
			},
		)

		users, err := repo.FindAll()

		assert.NoError(t, err)
		assert.Len(t, users, 2)
	})
}
