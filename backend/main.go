package main

import (
    "flag"
    "fmt"

    _ "github.com/jinzhu/gorm/dialects/mysql"
    "NagoBackend/config"
    "NagoBackend/database"
    "NagoBackend/server"
)

func main() {
    // コマンドライン引数「-e」の値を取得する（デフォルトは "development"）
    env := flag.String("e", "development", "")
    flag.Parse()

    config.Init(*env)
    database.Init()
    defer database.Close()
    if err := server.Init(); err != nil {
        fmt.Println("server error")
        panic(err)
    }
}
