package services

import (
	"errors"
	"train-http/internal/repositories"
	"train-http/internal/utils"
	"train-http/internal/validators"
)

type UserService interface {
	GetUser(userID string) (validators.GetUser, error)
	RegisterUser(user validators.PostUser) (string, error)

	//DeleteUser(userID string) (int, error)
}

type userService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) UserService {
	return &userService{repo: repo}
}

func (service *userService) GetUser(userID string) (validators.GetUser, error) {
	user, err := service.repo.GetUser(userID)

	if err != nil {
		return validators.GetUser{}, err
	}

	result := validators.GetUser{
		Username: user.Username,
		Email:    user.Email,
	}

	return result, nil
}

func (service *userService) RegisterUser(user validators.PostUser) (string, error) {
	checkUserExists := service.repo.CheckUserAlreadyExists(user.Email)

	if checkUserExists {
		return "", errors.New("User already exists")
	}
	userEmail, err := service.repo.RegisterUser(user)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(userEmail)
	if err != nil {
		return "", err
	}

	return token, nil

}

//func (service *userService) DeleteUser(userID string) (int, error) {
//	result, err := service.repo.DeleteUser(userID)
//
//	if err != nil {
//		return 0, err
//	}
//
//	return result, nil
//}
