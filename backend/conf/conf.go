package conf

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/spf13/viper"
)

// 設定型
type config struct {
    // yml形式に合わせて構造体で記述する
    Database struct {
        Host     string
        Port     string
        User     string
        Password string
        DBname   string
    }
    Server struct {
        Port string
    }
}

var C config

func ReadConf() {
    Conf := &C

    // 設定ファイル名を記載
    viper.SetConfigName("config")
    // ファイルタイプ
    viper.SetConfigType("yml")
    // ファイルパスの設定
    viper.AddConfigPath(filepath.Join("$GOPATH", "src", "NagoBackend", "conf"))
    // 環境変数から設定値を上書きできるように設定
    viper.AutomaticEnv()

    // conf読み取り
    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("config file read error")
        fmt.Println(err)
        os.Exit(1)
    }

    // UnmarshalしてCにマッピング
    if err := viper.Unmarshal(&Conf); err != nil {
        fmt.Println("config file Unmarshal error")
        fmt.Println(err)
        os.Exit(1)
    }
}
