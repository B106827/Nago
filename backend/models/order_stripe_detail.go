package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type OrderStripeDetail struct {
	ID         uint            `json:"-"  gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	OrderID    uint            `json:"-"  gorm:"column(order_id);not null;type(uint);"`
	StripeID   string          `json:"-"  gorm:"column(stripe_id);not null;type(varchar(255));"`
	ResultData datatypes.JSON  `json:"-"  gorm:"column(result_data);not null;type(json);"`
	CreatedAt  time.Time       `json:"-"`
	UpdatedAt  time.Time       `json:"-"`
}

func (OrderStripeDetail) TableName() string {
	return "order_stripe_detail"
}

func (osd *OrderStripeDetail) CreateInTx(tx *gorm.DB) error {
	return tx.Create(osd).Error
}
