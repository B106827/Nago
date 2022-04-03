package main

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "NagoBackend/conf"
    "NagoBackend/infrastructure/api/router"
    "NagoBackend/infrastructure/datastore"
    "NagoBackend/registry"
)

func main() {
    // conf読み取り
    conf.ReadConf()

    // DB取得
    conn := datastore.NewMySqlDB()

    // interactor
    r := registry.NewInteractor(conn)

    // 依存解決
    h := r.NewUserHandler()

    // echo
    e := echo.New()
    // middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // router
    router.NewRouter(e, h)

    // DB stop
    defer func () {
        if err := conn.Close(); err != nil {
            e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
        }
    }()
    // server start
    e.Logger.Fatal(e.Start(":" + conf.C.Server.Port))
}
