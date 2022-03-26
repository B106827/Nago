package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
    "log"
    "os"
)

func DBNew() *gorm.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASSWORD"),
            os.Getenv("DB_HOST"),
            os.Getenv("DB_PORT"),
            os.Getenv("DB_NAME"))

    config := gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    }
    db, err := gorm.Open(mysql.Open(dsn), &config)
    if err != nil {
        log.Println("dsn setting use .env files")
        log.Println("dsn: " + dsn)
        panic(err.Error())
    }
    return db
}
