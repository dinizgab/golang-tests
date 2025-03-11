package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/dinizgab/golang-tests/internal/models"
	"github.com/dinizgab/golang-tests/internal/repository"
	"github.com/dinizgab/golang-tests/internal/service"
)

type UserUsecase interface {
	FindAll() ([]models.User, error)
	FindByID(id string) (models.User, error)
	Save(user models.User) error
	FollowUser(userId string, followUserId string) error
}

type userUsecaseImpl struct {
	repo                repository.UserRepository
	notificationService service.NotificationService
}

func NewUserUsecase(
	repo repository.UserRepository,
	notificationService service.NotificationService,
) UserUsecase {
	return &userUsecaseImpl{
		repo,
		notificationService,
	}
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

func (u *userUsecaseImpl) FollowUser(followerUserId string, followedUserId string) error {
	// TODO - Implement pipeline pattern (good blog post)
	err := u.repo.FollowUser(followerUserId, followedUserId)
	if err != nil {
		return err
	}

	body, err := json.Marshal(map[string]string{
		"user_id":     followedUserId,
		"followed_by": followerUserId,
		"message":     "Hey, you have a new follower!",
	})
	if err != nil {
		return fmt.Errorf("UserUsecase.FollowUser: error marshalling notification body - %w", err)
	}

	err = u.notificationService.Publish(body)
	if err != nil {
		return fmt.Errorf("UserUsecase.FollowUser: error sending notification - %w", err)
	}

	return nil
}
