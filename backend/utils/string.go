package utils

import (
	"regexp"
)

// 氏名を正規表現で確認する
func CheckNameByRegexp(name string) ([][]string, string) {
	// 「(半角全角スペース以外の文字)(半角全角スペース)(文字)」
	r := regexp.MustCompile(`^([^ 　]+)( |　)+(.+)$`)
	if r.MatchString(name) {
		return r.FindAllStringSubmatch(name, -1), ""
	}
	return nil, "姓と名の間にスペースを入れてください"
}

// 郵便番号を正規表現で確認する
func CheckPostcodeByRegexp(postcode string) ([][]string, string) {
	// 「[半角全角3桁の数値][半角全角ハイフン][半角全角4桁の数値]」
	r := regexp.MustCompile(`^[0-9０-９]{3}[-ー][0-9０-９]{4}$`)
	if r.MatchString(postcode) {
		return r.FindAllStringSubmatch(postcode, -1), ""
	}
	return nil, "郵便番号の形式が不適切です"
}

// 市区町村以降の住所を正規表現で確認する
func CheckPrimaryAddressByRegexp(primaryAddress string) (string, string) {
	// 「(任意の文字+市区町村のいずれか)(半角全角混じりの任意文字)」
	r := regexp.MustCompile(`(.+?[市区町村])(.*[0-9０-９]+.*)`)
	if r.MatchString(primaryAddress) {
		return primaryAddress, ""
	}
	return "", "住所の形式が不適切です"
}

// 市区町村以降の住所を正規表現で確認する
func CheckPhoneNumberByRegexp(phoneNumber string) ([][]string, string) {
	// [半角全角0][半角全角1〜3桁の数字][半角全角ハイフン][半角全角1〜4桁の数字][半角全角4桁の数字]」
	r := regexp.MustCompile(`^[0０][0-9０-９]{1,3}[-ー][0-9０-９]{1,4}[-ー][0-9０-９]{4}$`)
	if r.MatchString(phoneNumber) {
		return r.FindAllStringSubmatch(phoneNumber, -1), ""
	}
	return nil, "電話番号の形式が不適切です"
}
