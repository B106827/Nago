package models

import (
	"NagoBackend/constants"
	"NagoBackend/database"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint   `json:"id"    gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	Email     string `json:"email" gorm:"column(email);size(255);unique;not null;type(varchar(255));"`
	Password  string `json:"-"     gorm:"column(password);size(255);not null;type(varchar(255));"`
	Name      string `json:"name"  gorm:"column(name);size(255);not null;type(varchar(255));"`
	Status    uint   `json:"-"     gorm:"column(status);default(1);type(int);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}

func (u *User) FindById(id int) (*User, error) {
	db := database.GetDB()
	var res User
	result := db.Where("id = ? AND status = ?", id, constants.USER_STATUS_REGISTERED).First(&res).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		// データが存在しない
		return nil, nil
	} else if result != nil {
		// 上記以外のエラー
		return nil, result
	}
	return &res, nil
}

func (u *User) FindByEmail(email string) (*User, error) {
	db := database.GetDB()
	var res User
	result := db.Where("email = ? AND status = ?", email, constants.USER_STATUS_REGISTERED).First(&res).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		// ユーザーが存在しない
		return nil, nil
	} else if result != nil {
		// 上記以外のエラー
		return nil, result
	}
	return &res, nil
}

func (u *User) Create() error {
	db := database.GetDB()
	return db.Create(u).Error
}
