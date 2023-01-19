package models

import (
	"NagoBackend/database"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type UserTemporary struct {
	ID        string    `json:"id"    gorm:"column(id);primaryKey;not null;type(char(10));"`
	Email     string    `json:"email" gorm:"column(email);size(255);type(varchar(255));"`
	ExpiredAt time.Time `json:"-"     gorm:"column(expired_at);not null;type(timestamp);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserTemporary) TableName() string {
	return "user_temporary"
}

func (ut *UserTemporary) FindById(id string) (*UserTemporary, error) {
	db := database.GetDB()
	var res UserTemporary
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

func (ut *UserTemporary) Create() error {
	db := database.GetDB()
	return db.Create(ut).Error
}

func (ut *UserTemporary) Delete() error {
	db := database.GetDB()
	return db.Delete(ut).Error
}
