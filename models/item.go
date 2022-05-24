package models

import "time"

type Item struct {
	ItemID      uint      `json:"item_id" gorm:"primaryKey"`
	ItemCode    string    `json:"item_code" gorm:"not null;unique;type:varchar(20)"`
	Description string    `json:"description" gorm:"not null;type:text"`
	Quantity    uint      `json:"quantity" gorm:"not null;type:uint"`
	OrderID     uint      `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
