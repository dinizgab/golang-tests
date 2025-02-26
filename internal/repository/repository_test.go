package repository

import (
	"testing"

	"github.com/google/uuid"
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

	t.Run("Test create new post for user", func(t *testing.T) {
		db.Exec("DELETE FROM users")
		db.Exec("DELETE FROM posts")

		user := models.User{
			ID:        uuid.New().String(),
			FirstName: "Gabriel",
			Username:  "dinizgab",
		}
		db.Exec("INSERT INTO users (id, first_name, username) VALUES ($1, $2, $3)", user.ID, user.FirstName, user.Username)

		post := models.Post{
			Title: "My first post",
			Body:  "This is my first post",
		}

		err = repo.SavePost(user.ID, post)
		assert.NoError(t, err)

		posts, err := repo.FindUserPosts(user.ID)

		assert.NoError(t, err)
		assert.Len(t, posts, 1)
		assert.Equal(t, "My first post", posts[0].Title)
		assert.Equal(t, "This is my first post", posts[0].Body)
	})

	t.Run("Test delete post", func(t *testing.T) {
		userId := uuid.New().String()
		postId := uuid.New().String()

		db.Exec("DELETE FROM users")
		db.Exec("DELETE FROM posts")
		db.Exec("INSERT INTO users (id, first_name, username) VALUES ($1, $2, $3)", userId, "Gabriel", "dinizgab")
		db.Exec("INSERT INTO posts (id, user_id, title, body) VALUES ($1, $2, $3, $4)", postId, userId, "My first post", "This is my first post")
    
        err := repo.DeletePost(postId)
        assert.NoError(t, err)

        posts, err := repo.FindUserPosts(userId)
        assert.NoError(t, err)
        assert.Len(t, posts, 0)
	})
}
