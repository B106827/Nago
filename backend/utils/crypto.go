package utils

import (
    "crypto/sha1"
    "crypto/rand"
    "encoding/hex"
)

func GetEncryptedHash(text string) string {
    textSha1 := sha1.Sum([]byte(text))
    return hex.EncodeToString(textSha1[:])
}

func GetRandomText(digits uint8) (string, error) {
    const keyText = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    // 乱数生成
    b := make([]byte, digits)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    // keyText からランダムに取り出してランダム文字列作成
    result := ""
    for _, v := range b {
        // index が keyText の長さに収まるように調整
        result += string(keyText[int(v)%len(keyText)])
    }
    return result, nil
}
