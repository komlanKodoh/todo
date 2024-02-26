package repository

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	Id          uint   `gorm:"primarykey" json:"Id"`
	Title       string `gorm:"type:varchar(100);"`
	Description string `gorm:"type:varchar(1000);"`

	CompletedAt gorm.DeletedAt
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
