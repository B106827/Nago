package models

import (
    "NagoBackend/database"
    "time"
)

type UserTemporary struct {
	ID           string    `json:"id"           gorm:"column(id);primaryKey;not null;type(char(10));"`
	Email        string    `json:"email"        gorm:"column(email);size(255);type(varchar(255));"`
	UserID       uint      `json:"-"            gorm:"column(user_id);type(int);"`
	RedirectPath string    `json:"redirectPath" gorm:"column(redirect_path);size(255);type(varchar(255));"`
    ExpiredAt    time.Time `json:"-"            gorm:"column(expired_at);not null;type(datetime);"`
    CreatedAt    time.Time `json:"-"            gorm:"column(created_at);not null;type(datetime);"`
    UpdatedAt    time.Time `json:"-"            gorm:"column(updated_at);not null;type(datetime);"`
}

func (UserTemporary) TableName() string {
    return "user_temporary"
}

// find user_temporary by id
func (ut *UserTemporary) FindById(id string) error {
    db := database.GetDB()
    return db.Where("id = ?", id).First(ut).Error
}

func (ut *UserTemporary) Create() error {
    db := database.GetDB()
    return db.Create(ut).Error
}

func (ut *UserTemporary) Delete() error {
    db := database.GetDB()
    return db.Delete(ut).Error
}
