package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OrderDetail struct {
	ID         uint      `json:"id"       gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	OrderID    uint      `json:"-"        gorm:"column(order_id);not null;type(uint);"`
	ProductID  uint      `json:"-"        gorm:"column(product_id);not null;type(uint);"`
	Price      uint      `json:"price"    gorm:"column(price);not null;type(uint);"`
	Num        uint      `json:"num"      gorm:"column(num);not null;type(uint);"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	Product    Product   `json:"product"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}

func (od *OrderDetail) CreateInTx(tx *gorm.DB, orderId uint, params *[]map[string]uint) error {
	sql := "INSERT INTO order_detail (order_id, product_id, price, num) VALUES "
	var values []interface{}
	for i, param := range *params {
		sql += "(" + GetParamMarks(len(param)) + ")"
		if i != len(*params)-1 {
			sql += ","
		}
		values = append(values, orderId, param["productId"], param["price"], param["num"]) 
	}
	return tx.Exec(sql, values...).Error
}
