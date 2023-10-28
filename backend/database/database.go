package database

import (
	"NagoBackend/config"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var d *gorm.DB

func Init() {
	e := echo.New()
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
		e.Logger.Fatal(err)
	}
	// 応答確認
	err = d.DB().Ping()
	if err != nil {
		fmt.Println("db no response")
		e.Logger.Fatal(err)
	}
	// SQLログ出力
	d.LogMode(true)
}

func GetDB() *gorm.DB {
	return d
}

// トランザクション処理
func ExecuteInTx(fn func(*gorm.DB) error) error {
	db := GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func Close() {
	d.Close()
}
