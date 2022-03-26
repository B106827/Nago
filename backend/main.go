package main

import (
    "fmt"
    "NagoBackend/router"
    "os"
)

func main() {
    fmt.Pritln("Server starting...")
    e := router.New()
    e.Start(":" + os.Getenv("PORT"))
}
