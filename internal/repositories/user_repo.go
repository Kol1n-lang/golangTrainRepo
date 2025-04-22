package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"train-http/internal/models"
	"train-http/internal/utils"
	"train-http/internal/validators"
)

type UserRepoInterface interface {
	GetUser(userID string) (models.User, error)
	RegisterUser(userData validators.PostUser) (string, error)
	CheckUserAlreadyExists(email string) bool
	//DeleteUser(userID string) (int, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &userRepo{db: db}
}

func (repo *userRepo) GetUser(userID string) (models.User, error) {
	var user models.User
	err := repo.db.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (repo *userRepo) CheckUserAlreadyExists(email string) bool {
	//var user models.User
	log.Print(email)
	return false
}

func (repo *userRepo) RegisterUser(userData validators.PostUser) (string, error) {
	var user models.User
	hashedPassword, err := utils.HashedPassword(userData.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword
	user.Email = userData.Email
	user.Username = userData.Username
	if err := repo.db.Create(&user).Error; err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}
	return user.Email, nil
}

//
//func (repo *userRepo) DeleteUser(userID string) (int, error) {
//
//}
