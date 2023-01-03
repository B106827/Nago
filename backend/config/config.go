package config

import (
    "fmt"
    "github.com/spf13/viper"
)

var c *viper.Viper

func Init(env string) {
    c = viper.New()
    c.SetConfigFile("yaml")
    c.SetConfigName(env)
    c.AddConfigPath("config/environments/")
    if err := c.ReadInConfig(); err != nil {
        fmt.Println("config file read error")
        panic(err)
    }
    c.SetConfigName("common")
    c.AddConfigPath("config/common/")
    if err := c.MergeInConfig(); err != nil {
        fmt.Println("config file read error")
        panic(err)
    }
    c.SetConfigName("private")
    c.AddConfigPath("config/common/")
    if err := c.MergeInConfig(); err != nil {
        fmt.Println("config file read error")
        panic(err)
    }
}

func GetConfig() *viper.Viper {
    return c
}
