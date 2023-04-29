package models

import (
	"NagoBackend/constants"
	"NagoBackend/database"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	ID        uint      `json:"id"       gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	UserID    uint      `json:"-"        gorm:"column(user_id);not null;type(uint);"`
	ProductID uint      `json:"-"        gorm:"column(product_id);not null;type(uint);"`
	Num       uint      `json:"num"      gorm:"column(num);not null;type(uint);"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Product   Product   `json:"product"`
}

func (Cart) TableName() string {
	return "cart"
}

func (c *Cart) FindById(id uint) (*Cart, error) {
	db := database.GetDB()
	var res Cart
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

func (c *Cart) FindByUserIdAndProductId(userId uint, productId uint) (*Cart, error) {
	db := database.GetDB()
	var res Cart
	result := db.Where("user_id = ? AND product_id = ?", userId, productId).First(&res).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		// データが存在しない
		return nil, nil
	} else if result != nil {
		// 上記以外のエラー
		return nil, result
	}
	return &res, nil
}

func (c *Cart) FindByUserIdWithProduct(userId uint) ([]Cart, error) {
	db := database.GetDB()
	var res []Cart
	result := db.Preload("Product", "status = ?", constants.PRODUCT_STATUS_ACTIVE).Preload("Product.Images").Where("user_id = ?", userId).Find(&res).Error
	if result != nil {
		// エラーが発生した場合
		return nil, result
	} else if len(res) == 0 {
		// 上記以外のエラー
		return nil, nil
	}
	return res, nil
}

func (c *Cart) Create() error {
	db := database.GetDB()
	return db.Create(c).Error
}

func (c *Cart) Update(num uint) error {
	db := database.GetDB()
	return db.Model(c).Update("num", num).Error
}

func (c *Cart) Delete() error {
	db := database.GetDB()
	return db.Delete(c).Error
}

func (c *Cart) DeleteCartDataInTx(tx *gorm.DB, userId uint) error {
	return tx.Model(c).Where("user_id = ?", userId).Delete(c).Error
}
