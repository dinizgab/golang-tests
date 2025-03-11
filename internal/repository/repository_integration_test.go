package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dinizgab/golang-tests/internal/config"
	"github.com/dinizgab/golang-tests/internal/db"
	"github.com/dinizgab/golang-tests/internal/models"
)

func TestRepositoryIntegration(t *testing.T) {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		t.Fatal(err)
	}
	db, err := db.New(dbConfig)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepository(db)

    db.Exec("DELETE FROM users")
    db.Exec("DELETE FROM posts")

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

		db.Exec("INSERT INTO users (first_name, username) VALUES ($1, $2)", "Gabriel", "dinizgab")
		db.Exec("INSERT INTO users (first_name, username) VALUES ($1, $2)", "Joao", "john")

		users, err := repo.FindAll()

		assert.NoError(t, err)
		assert.Len(t, users, 2)
	})

	t.Run("Test find user by id", func(t *testing.T) {
		db.Exec("DELETE FROM users")

		user := models.User{
			ID:        uuid.New().String(),
			FirstName: "Gabriel",
			Username:  "dinizgab",
		}

		db.Exec("INSERT INTO users (id, first_name, username) VALUES ($1, $2, $3)", user.ID, user.FirstName, user.Username)

		user, err := repo.FindByID(user.ID)

		assert.NoError(t, err)
		assert.Equal(t, "Gabriel", user.FirstName)
		assert.Equal(t, "dinizgab", user.Username)
	})
}
