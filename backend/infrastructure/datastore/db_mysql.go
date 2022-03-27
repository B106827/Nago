package datastore

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "NagoBackend/conf"
)

func NewMySqlDB() *gorm.DB {
    connString := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        conf.C.Database.User,
        conf.C.Database.Password,
        conf.C.Database.Host,
        conf.C.Database.Port,
        conf.C.Database.DBname,
    )

    conn, err := gorm.Open("mysql", connString)

    if err != nil {
        fmt.Println("db connection error")
        panic(err)
    }

    // 応答確認
    err = conn.DB().Ping()
    if err != nil {
        fmt.Println("db no response")
        panic(err)
    }
    // SQLログの詳細出力
    conn.LogMode(true)
    // DBエンジンを設定
    return conn
}
