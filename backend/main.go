package main

import (
	"flag"
	"fmt"

	"NagoBackend/config"
	"NagoBackend/database"
	"NagoBackend/server"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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
