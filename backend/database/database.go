package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "NagoBackend/config"
)

var d *gorm.DB

func Init() {
    c := config.GetConfig()
    var err error
    connString := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        c.GetString("db.user"),
        c.GetString("db.password"),
        c.GetString("db.host"),
        c.GetString("db.port"),
        c.GetString("db.dbname"),
    )
    fmt.Sprintf("%s", connString)
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
