package config

import (
    "fmt"
    "github.com/spf13/viper"
)

var c *viper.Viper

func Init(env string) {
    c = viper.New()
    c.SetConfigFile("config/common/config.yml")
    c.SetConfigFile("yaml")
    c.SetConfigName("config")
    c.AddConfigPath("config/environments/" + env)
    if err := c.ReadInConfig(); err != nil {
        fmt.Println("config file read error")
        panic(err)
    }
}

func GetConfig() *viper.Viper {
    return c
}
