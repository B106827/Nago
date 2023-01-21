package database

import (
	"NagoBackend/config"
	"fmt"

	"github.com/jinzhu/gorm"
)

var d *gorm.DB

func Init() {
	c := config.GetConfig()
	var err error
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.GetString("db.user"),
		c.GetString("db.password"),
		c.GetString("db.host"),
		c.GetInt("db.port"),
		c.GetString("db.dbname"),
	)
	d, err = gorm.Open(c.GetString("db.provider"), connString)
	if err != nil {
		fmt.Println("db connection error")
		panic(err)
	}
	// 応答確認
	err = d.DB().Ping()
	if err != nil {
		fmt.Println("db no response")
		panic(err)
	}
	// SQLログ出力
	d.LogMode(true)
}

func GetDB() *gorm.DB {
	return d
}

func Close() {
	d.Close()
}
