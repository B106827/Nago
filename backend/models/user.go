package models

import (
    "NagoBackend/constants"
    "NagoBackend/database"
    "time"
)

type User struct {
	ID         uint      `json:"id"    gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	Email      string    `json:"email" gorm:"column(email);size(255);unique;not null;type(varchar(255));"`
	Password   string    `json:"-"     gorm:"column(password);size(255);not null;type(varchar(255));"`
	LineUserID string    `json:"-"     gorm:"column(line_user_id);size(255);type(varchar(255));"`
	Tel        string    `json:"-"     gorm:"column(tel);size(13);type(varchar(13));"`
	Name       string    `json:"name"  gorm:"column(name);size(255);type(varchar(255));"`
	Address    string    `json:"-"     gorm:"column(address);size(255);type(varchar(255));"`
	Status     uint      `json:"-"     gorm:"column(status);default(1);type(int);"`
    CreatedAt  time.Time `json:"-"     gorm:"column(created_at);not null;type(datetime);"`
    UpdatedAt  time.Time `json:"-"     gorm:"column(updated_at);not null;type(datetime);"`
}

func (User) TableName() string {
    return "user"
}

// find active user by email
func (u *User) FindByEmail(email string) error {
    db := database.GetDB()
    return db.Where("email = ? AND status = ?", email, constants.USER_STATUS_REGISTERED).First(u).Error
}

func (u *User) Create() error {
    db := database.GetDB()
    return db.Create(u).Error
}
