package models

import (
	"time"
)

type ProductImage struct {
	ID        uint   `json:"id"  gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	ProductID uint   `json:"-"   gorm:"column(product_id);not null;type(int);"`
	Url       string `json:"url" gorm:"column(url);type(text);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ProductImage) TableName() string {
	return "product_image"
}
