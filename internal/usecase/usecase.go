package usecase

import (
	"github.com/dinizgab/golang-tests/internal/models"
	"github.com/dinizgab/golang-tests/internal/repository"
)


type UserUsecase interface {
	FindAll() ([]models.User, error)
	FindByID(id string) (models.User, error)
	Save(user models.User) error
	SavePost(userId string, post models.Post) error
    FindUserPosts(userId string) ([]models.Post, error)
    DeletePost(postId string) error
    FollowUser(userId string, followUserId string) error
}

type userUsecaseImpl struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{repo}
}

func (u *userUsecaseImpl) FindAll() ([]models.User, error) {
	return u.repo.FindAll()
}

func (u *userUsecaseImpl) FindByID(id string) (models.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecaseImpl) Save(user models.User) error {
	return u.repo.Save(user)
}

func (u *userUsecaseImpl) SavePost(userId string, post models.Post) error {
	return u.repo.SavePost(userId, post)
}

func (u *userUsecaseImpl) FindUserPosts(userId string) ([]models.Post, error) {
	return u.repo.FindUserPosts(userId)
}

func (u *userUsecaseImpl) DeletePost(postId string) error {
	return u.repo.DeletePost(postId)
}

func (u *userUsecaseImpl) FollowUser(userId string, followUserId string) error {
	// TODO - Add the call to the broker to notify the user that he is being followed

	return u.repo.FollowUser(userId, followUserId)
}
