package tokens

import (
    "math/big"
    "crypto/rand"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var AuthTokenLength = 10

func GenerateToken() (string, error) {
    var err error
    var num *big.Int
    b := make([]rune, AuthTokenLength)
    for i := range b {
        num, err = rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
        if err != nil {
            return "", err
        }
        b[i] = chars[num.Int64()]
    }
    return string(b), nil
}

