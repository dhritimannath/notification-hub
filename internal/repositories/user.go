package repositories

import (
	"fmt"

	"github.com/dhritimannath/notification-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(*models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(uuid.UUID) (models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&models.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}


func (u UserRepositoryImpl) Create(user *models.User) (models.User, error) {
	var err = u.db.Create(user).Error
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (u UserRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	var err = u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepositoryImpl) FindById(id uuid.UUID) (models.User, error) {
	user  :=  models.User{
		ID: id,
	}
	var err = u.db.First(&user).Error
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	return user, nil
}