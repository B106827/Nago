package models

import (
    "NagoBackend/database"
    "time"
)

type User struct {
	ID         uint      `json:"id"    gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	Email      string    `json:"email" gorm:"column(email);size(255);unique;not null;type(varchar(255));"`
	Password   string    `json:"-"     gorm:"column(password);size(255);not null;type(varchar(255));"`
	LineUserID *string   `json:"-"     gorm:"column(line_user_id);size(255);type(varchar(255));"`
	Tel        *string   `json:"-"     gorm:"column(tel);size(13);type(varchar(13));"`
	Name       *string   `json:"name"  gorm:"column(name);size(255);type(varchar(255));"`
	Address    *string   `json:"-"     gorm:"column(address);size(255);type(varchar(255));"`
	Status     uint      `json:"-"     gorm:"column(status);default(1);type(int);"`
	CreatedOn  time.Time `json:"-"     gorm:"column(created_on);autoCreateTime;type(datetime);"`
	ModifiedOn time.Time `json:"-"     gorm:"column(modified_on);autoUpdateTime;type(timestamp);"`
}

func (User) TableName() string {
    return "user"
}

// create a user
func (u *User) Create() (err error) {
    db := database.GetDB()
    return db.Create(u).Error
}

// find user by id
func (u *User) FindByID(id uint) (err error) {
    db := database.GetDB()
    return db.Where("id = ?", id).First(u).Error
}

// find user by email
func (u *User) FindByEmail(email string) (err error) {
    db := database.GetDB()
    return db.Where("email = ?", email).First(u).Error
}
