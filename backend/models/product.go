package models

import (
	"NagoBackend/constants"
	"NagoBackend/database"
	"time"
)

type Product struct {
	ID          uint           `json:"id"          gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	Name        string         `json:"name"        gorm:"column(title);size(255);not null;type(varchar(255));"`
	SubTitle    string         `json:"subTitle"    gorm:"column(sub_title);size(255);type(varchar(255));"`
	Price       uint           `json:"price"       gorm:"column(price);not null;type(int);"`
	Description string         `json:"description" gorm:"column(description);type(text);"`
	Status      uint           `json:"-"           gorm:"column(status);default(1);type(int);"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	Images      []ProductImage `json:"images"`
}

func (Product) TableName() string {
	return "product"
}

func (p *Product) GetActiveProductsWithImages() ([]Product, error) {
	db := database.GetDB()
	var res []Product
	result := db.Preload("Images").Where("product.status = ?", constants.PRODUCT_STATUS_ACTIVE).Find(&res).Error
	if result != nil {
		// エラーが発生した場合
		return nil, result
	} else if len(res) == 0 {
		// 商品が一件もない場合
		return nil, nil
	}
	return res, nil
}
