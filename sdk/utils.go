package sdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

var (
	// Debug Debug Mode
	Debug bool = false
)

// SetDebug Set Debug Mode
func SetDebug(b bool) {
	Debug = b
}

func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

func unPadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

func aesDecrypt(message, key, iv string) (data []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	cipherText := []byte(message)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	return unPadding(plainText), nil
}
