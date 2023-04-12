package forms

import (
	"NagoBackend/utils"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type OrderDeliveryInfoForm struct {
	Total            uint   `json:"total"`
	Name             string `json:"name"`
	Postcode         string `json:"postcode"`
	PrefId           uint   `json:"prefId"`
	PrimaryAddress   string `json:"primaryAddress"`
	SecondaryAddress string `json:"secondaryAddress"`
	PhoneNumber      string `json:"phoneNumber"`
}

func (f OrderDeliveryInfoForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(
			&f.Total,
			validation.Required.Error("パラメータが不正です"),
			validation.By(func(value interface{}) error {
				total, ok := value.(uint)
				if !ok {
					return validation.NewError("custom", "パラメータが不正です")
				}
				// uint→int→string に変換して正規表現のチェックを通す
				s := strconv.Itoa(int(total))
				if !utils.CheckNumericByRegexp(s) {
					return validation.NewError("custom", "パラメータが不正です")
				}
				return nil
			}),
		),
		validation.Field(
			&f.Name,
			validation.Required.Error("お名前を入力してください"),
			validation.By(func(value interface{}) error {
				name, ok := value.(string)
				if !ok {
					return validation.NewError("custom", "パラメータが不正です")
				}
				if _, errMsg := utils.CheckNameByRegexp(name); errMsg != "" {
					return validation.NewError("custom", errMsg)
				}
				return nil
			}),
		),
		validation.Field(
			&f.Postcode,
			validation.Required.Error("郵便番号を入力してください"),
			validation.By(func(value interface{}) error {
				postcode, ok := value.(string)
				if !ok {
					return validation.NewError("custom", "パラメータが不正です")
				}
				if _, errMsg := utils.CheckPostcodeByRegexp(postcode); errMsg != "" {
					return validation.NewError("custom", errMsg)
				}
				return nil
			}),
		),
		validation.Field(
			&f.PrefId,
			validation.Required.Error("都道府県を選択してください"),
		),
		validation.Field(
			&f.PrimaryAddress,
			validation.Required.Error("住所1を入力してください"),
			validation.By(func(value interface{}) error {
				primaryAddress, ok := value.(string)
				if !ok {
					return validation.NewError("custom", "パラメータが不正です")
				}
				if _, errMsg := utils.CheckPrimaryAddressByRegexp(primaryAddress); errMsg != "" {
					return validation.NewError("custom", errMsg)
				}
				return nil
			}),
		),
		validation.Field(
			&f.PhoneNumber,
			validation.Required.Error("電話番号を入力してください"),
			validation.By(func(value interface{}) error {
				phoneNumber, ok := value.(string)
				if !ok {
					return validation.NewError("custom", "パラメータが不正です")
				}
				if _, errMsg := utils.CheckPhoneNumberByRegexp(phoneNumber); errMsg != "" {
					return validation.NewError("custom", errMsg)
				}
				return nil
			}),
		),
	)
}
