package utils

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

func GetEncryptedText(text string) (string, error) {
    // 暗号化したい文字列
    plainText := []byte(text)
    // aes の暗号化文字列
    key := []byte("abcdefg12798akljzmknm.ahkjkljl;k")

    // 暗号化アルゴリズムAESを作成
    c, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    // IV（ストリームの初期値)を作成
    cipherText := make([]byte, aes.BlockSize+len(plainText))
    iv := cipherText[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    // 暗号化
    encryptStream := cipher.NewCTR(c, iv)
    encryptStream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
    return string(cipherText), nil
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
