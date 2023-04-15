package models

import (
	"NagoBackend/database"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
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

const (
	STATUS_PENDING = 0
)

func (Order) TableName() string {
	return "order"
}

func (o *Order) FindById(id uint) (*Order, error) {
	db := database.GetDB()
	var res Order
	result := db.Where("id = ?", id).First(&res).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		// データが存在しない
		return nil, nil
	} else if result != nil {
		// 上記以外のエラー
		return nil, result
	}
	return &res, nil
}

func (o *Order) Create(tx *gorm.DB) error {
	return tx.Create(o).Error
}
