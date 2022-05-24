package models

import "time"

type Order struct {
	OrderID      uint      `json:"order_id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name" gorm:"not null;type:varchar(50)"`
	OrderedAt    string    `json:"ordered_at" gorm:"not null;type:varchar(100)"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
