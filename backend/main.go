package main

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "NagoBackend/handler"
)

func main() {
    fmt.Println("Server starting...")
    e := echo.New()
    e.GET("/", handler.HandleGetPosts)
    e.Start(":8081")

    // regist handler
    //group := e.Group("/api")
    //msg_handler := controller.MsgHandlerFactory()
    //for uri, instance := range msg_handler {
    //    group.GET(uri, instance.Get)
    //    group.POST(uri, instance.Post)
    //    group.PUT(uri, instance.Put)
    //    group.DELETE(uri, instance.Delete)
    //}
}
