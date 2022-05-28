package models

import "time"

type Order struct {
	OrderID      uint      `json:"order_id,omitempty" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name" gorm:"not null;type:varchar(50)"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
