package repository

import (
	"testing"

	"github.com/dinizgab/golang-tests/internal/db"
	"github.com/dinizgab/golang-tests/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryUnit(t *testing.T) {
	queryInsertUser := `INSERT INTO users (first_name, username) VALUES ($1, $2)`

	database := new(db.DatabaseMock)
	repo := NewUserRepository(database)

	t.Run("Test create new user", func(t *testing.T) {
		res := new(db.SqlResultMock)
		database.On("Exec", queryInsertUser, []interface{}{"gabriel", "dinizgab"}).Return(res, nil)

		err := repo.Save(models.User{
			FirstName: "gabriel",
			Username:  "dinizgab",
		})

		assert.NoError(t, err)
		database.AssertExpectations(t)
	})
}
