package models

import (
	"NagoBackend/database"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	ID         uint          `json:"id"         gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	UserID     uint          `json:"-"          gorm:"column(user_id);not null;type(uint);"`
	TotalPrice uint          `json:"totalPrice" gorm:"column(total_price);not null;type(uint);"`
	Status     uint          `json:"-"          gorm:"column(status);not null;default(0);type(uint);"`
	OrderedAt  time.Time     `json:"orderedAt"  gorm:"column(ordered_at);not null;type(timestamp);"`
	CreatedAt  time.Time     `json:"-"`
	UpdatedAt  time.Time     `json:"-"`
	Details    []OrderDetail `json:"details"`
}

const (
	ORDER_STATUS_PENDING  = 0
	ORDER_STATUS_COMPLETE = 1
	ORDER_STATUS_CANCEL   = 2
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

func (o *Order) CreateInTx(tx *gorm.DB) error {
	return tx.Create(o).Error
}

func (o *Order) UpdateStatusCompleteInTx(tx *gorm.DB) error {
	return tx.Model(o).Update("status", ORDER_STATUS_COMPLETE).Error
}

func (o *Order) UpdateStatusCancel() error {
	db := database.GetDB()
	return db.Model(o).Update("status", ORDER_STATUS_CANCEL).Error
}

// 購入履歴
func (o *Order) FindByUserIdWithDetail(userId uint) ([]Order, error) {
	db := database.GetDB()
	var res []Order
	result := db.Preload("Details").Preload("Details.Product").Preload("Details.Product.Images").Where("user_id = ? AND status = ?", userId, ORDER_STATUS_COMPLETE).Find(&res).Error
	if result != nil {
		// データが発生した場合
		return nil, result
	} else if len(res) == 0 {
		// 履歴が一件もない場合
		return nil, nil
	}
	return res, nil
}
