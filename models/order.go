package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderId      uint           `gorm:"primaryKey; not null" json:"order_id"`
	CustomerName string         `gorm:"not_null; VARCHAR(255)" json:"customer_name"`
	Items        []Item         `gorm:"foreignKey:OrderId; constraint:OnDelete:CASCADE" json:"items"`
	OrderedAt    time.Time      `json:"ordered_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
