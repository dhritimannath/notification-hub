package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID uuid.UUID `gorm:"primary_key; not null" json:"id"`
	UserId uuid.UUID `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Token string `gorm:"type:text; not null" json:"token"`
	Type string `gorm:"type:varchar(255); not null" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}