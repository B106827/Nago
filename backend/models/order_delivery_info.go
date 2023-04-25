package models

import (
	forms "NagoBackend/forms/order"
	"NagoBackend/utils"
	"time"

	"github.com/jinzhu/gorm"
)

type OrderDeliveryInfo struct {
	ID                   uint      `json:"id" gorm:"column(id);primaryKey;autoIncrement;not null;type(uint);"`
	OrderID              uint      `json:"-"  gorm:"column(order_id);not null;type(uint);"`
	LastName             string    `json:"-"  gorm:"column(last_name);not null;type(varchar(255));"`
	FirstName            string    `json:"-"  gorm:"column(first_name);not null;type(varchar(255));"`
	PrimaryPostcode      string    `json:"-"  gorm:"column(primary_postcode);not null;type(varchar(3));"`
	SecondaryPostcode    string    `json:"-"  gorm:"column(secondary_postcode);not null;type(varchar(4));"`
	PrefID               uint      `json:"-"  gorm:"column(pref_id);not null;type(uint);"`
	PrimaryAddress       string    `json:"-"  gorm:"column(primary_address);not null;type(varchar(255));"`
	SecondaryAddress     *string   `json:"-"  gorm:"column(secondary_address);type(varchar(255));"`
	PrimaryPhoneNumber   string    `json:"-"  gorm:"column(primary_phone_number);not null;type(varchar(4));"`
	SecondaryPhoneNumber string    `json:"-"  gorm:"column(secondary_phone_number);not null;type(varchar(4));"`
	ThirdPhoneNumber     string    `json:"-"  gorm:"column(third_phone_number);not null;type(varchar(4));"`
	CreatedAt            time.Time `json:"-"`
	UpdatedAt            time.Time `json:"-"`
}

func (OrderDeliveryInfo) TableName() string {
	return "order_delivery_info"
}

func (odi *OrderDeliveryInfo) CreateInTx(tx *gorm.DB, orderId uint, odif forms.OrderDeliveryInfoForm) error {
	name, _               := utils.CheckNameByRegexp(odif.Name)
	postcode, _           := utils.CheckPostcodeByRegexp(odif.Postcode)
	phoneNumber, _        := utils.CheckPhoneNumberByRegexp(odif.PhoneNumber)
	odi.OrderID            = orderId
	odi.LastName           = name[0][1]
	odi.FirstName          = name[0][2]
	odi.PrimaryPostcode    = utils.ConvertHankakuAlphanumeric(postcode[0][1])
	odi.SecondaryPostcode  = utils.ConvertHankakuAlphanumeric(postcode[0][2])
	odi.PrefID             = odif.PrefId
	odi.PrimaryAddress     = odif.PrimaryAddress
	if odif.SecondaryAddress != "" {
		odi.SecondaryAddress = &odif.SecondaryAddress
	} else {
		// 住所2が入力されていない場合はNULLにしておく
		odi.SecondaryAddress = nil
	}
	odi.PrimaryPhoneNumber   = utils.ConvertHankakuAlphanumeric(phoneNumber[0][1])
	odi.SecondaryPhoneNumber = utils.ConvertHankakuAlphanumeric(phoneNumber[0][2])
	odi.ThirdPhoneNumber     = utils.ConvertHankakuAlphanumeric(phoneNumber[0][3])

	return tx.Create(odi).Error
}
