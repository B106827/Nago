package models

import (
	"NagoBackend/database"
	"time"
)

type PrefMaster struct {
	ID        uint      `json:"id"   gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	Name      string    `json:"name" gorm:"column(title);size(255);not null;type(varchar(255));"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (PrefMaster) TableName() string {
	return "pref_master"
}

func (pm *PrefMaster) GetPrefMaster() ([]PrefMaster, error) {
	db := database.GetDB()
	var res []PrefMaster
	result := db.Find(&res).Error
	if result != nil {
		// エラーが発生した場合
		return nil, result
	} else if len(res) == 0 {
		// 商品が一件もない場合
		return nil, nil
	}
	return res, nil
}
