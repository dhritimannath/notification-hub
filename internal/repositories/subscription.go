package repositories

import (
	"fmt"

	"github.com/dhritimannath/notification-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	Create(*models.Subscription) (models.Subscription, error)
	FindAllByUserId(uuid.UUID) ([]models.Subscription, error)
	FindById(uuid.UUID) (models.Subscription, error)
	DeleteById(uuid.UUID) error
	Save(*models.Subscription) (models.Subscription, error)
}

type SubscriptionRepositoryImpl struct {
	db *gorm.DB
}

func SubscriptionRepositoryInit(db *gorm.DB) *SubscriptionRepositoryImpl {
	db.AutoMigrate(&models.Subscription{})
	return &SubscriptionRepositoryImpl{
		db: db,
	}
}


func (s SubscriptionRepositoryImpl) Create(subscription *models.Subscription) (models.Subscription, error) {
	var err = s.db.Create(subscription).Error
	if err != nil {
		return models.Subscription{}, err
	}
	return *subscription, nil
}

func (s SubscriptionRepositoryImpl) Save(subscription *models.Subscription) (models.Subscription, error) {
	var err = s.db.Save(subscription).Error
	if err != nil {
		return models.Subscription{}, err
	}
	return *subscription, nil
}

func (s SubscriptionRepositoryImpl) FindById(id uuid.UUID) (models.Subscription, error) {
	subscription :=  models.Subscription{
		ID: id,
	}
	var err = s.db.First(&subscription).Error
	if err != nil {
		return models.Subscription{}, err
	}
	return subscription, nil
}


func (s SubscriptionRepositoryImpl) FindAllByUserId(userId uuid.UUID) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	var err = s.db.Where(&models.Subscription{UserId: userId}).Find(&subscriptions).Error
	if err != nil {
		fmt.Println(err)
		return []models.Subscription{}, err
	}
	return subscriptions, nil
}

func (s SubscriptionRepositoryImpl) DeleteById(id uuid.UUID) error {
	var result = s.db.Delete(&models.Subscription{}, id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}