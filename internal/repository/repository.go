package repository

import (
	"database/sql"
	"fmt"

	"github.com/dinizgab/golang-tests/internal/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id string) (models.User, error)
	Save(user models.User) error
	SavePost(userId string, post models.Post) error
    FindUserPosts(userId string) ([]models.Post, error)
    DeletePost(postId string) error
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

func (r *userRepositoryImpl) FindByID(id string) (models.User, error) {
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

func (r *userRepositoryImpl) SavePost(userId string, post models.Post) error {
	query := `INSERT INTO posts (user_id, title, body) VALUES ($1, $2, $3)`

	_, err := r.db.Exec(query, userId, post.Title, post.Body)
	if err != nil {
		return fmt.Errorf("UserRepository.SavePost: error saving post - %w", err)
	}

	return nil
}

func (r *userRepositoryImpl) FindUserPosts(userId string) ([]models.Post, error) {
    posts := []models.Post{}
	query := `SELECT id, title, body FROM posts WHERE user_id = $1`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.FindUserPosts: error fetching posts - %w", err)
	}

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			return nil, fmt.Errorf("UserRepository.FindUserPosts: error scanning posts - %w", err)
		}

		posts = append(posts, post)
	}

    return posts, nil
}

func (r *userRepositoryImpl) DeletePost(postId string) error {
    query := `DELETE FROM posts WHERE id = $1`

    _, err := r.db.Exec(query, postId)
    if err != nil {
        return fmt.Errorf("UserRepository.DeletePost: error deleting post - %w", err)
    }

    return nil
}
