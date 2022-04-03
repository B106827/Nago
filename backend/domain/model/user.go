package model

type User struct {
    ID   int    `gorm:"primary_key"`
    Name string `"name"`
}
