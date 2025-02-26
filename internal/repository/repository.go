package repository

import (
	"database/sql"
	"fmt"

	"github.com/dinizgab/golang-tests/internal/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id int) (models.User, error)
	Save(user models.User) error
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) FindAll() ([]models.User, error) {
    var users []models.User
    query := `SELECT id, first_name, username FROM users`

    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("UserRepository.FindAll: error fetching users - %w", err)
    }

    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.FirstName, &user.Username)
        if err != nil {
            return nil, fmt.Errorf("UserRepository.FindAll: error scanning users - %w", err)
        }

        users = append(users, user)
    }

	return users, nil
}

func (r *userRepositoryImpl) FindByID(id int) (models.User, error) {
	var user models.User
	query := `SELECT id, first_name, username FROM users WHERE id = $1`

	row := r.db.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.FirstName, &user.Username)
	if err != nil {
		return models.User{}, fmt.Errorf("UserRepository.FindByID: error fetching user - %w", err)
	}

	return user, nil
}

func (r *userRepositoryImpl) Save(user models.User) error {
	query := `INSERT INTO users (first_name, username) VALUES ($1, $2)`

	_, err := r.db.Exec(query, user.FirstName, user.Username)
	if err != nil {
		return fmt.Errorf("UserRepository.Save: error saving user - %w", err)
	}

	return nil
}
