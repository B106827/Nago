package models

import (
	"NagoBackend/database"
	"time"
)

type Order struct {
	ID         uint      `json:"id" gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	UserID     uint      `json:"-"  gorm:"column(user_id);not null;type(uint);"`
	TotalPrice uint      `json:"-"  gorm:"column(total_price);not null;type(uint);"`
	Status     uint      `json:"-"  gorm:"column(status);not null;default(0);type(uint);"`
	OrderedAt  time.Time `json:"-"  gorm:"column(ordered_at);not null;type(timestamp);"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

func (Order) TableName() string {
	return "order"
}

func (o *Order) Create() error {
	db := database.GetDB()
	return db.Create(o).Error
}
