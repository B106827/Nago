package model

import (
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
	CreatedOn  time.Time `json:"-"     gorm:"column(created_on);autoCreateTime;type(datetime);"`
	ModifiedOn time.Time `json:"-"     gorm:"column(modified_on);autoUpdateTime;type(timestamp);"`
}

func (User) TableName() string {
    return "user"
}
