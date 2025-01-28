package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func encrypt(plaintext string, key []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func decrypt(encrypted string, key []byte) string {
	ciphertext, _ := base64.StdEncoding.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func main() {
	key := []byte("mysecretencryptionkey") // 16, 24, or 32 bytes
	plaintext := "Hello, AES Encryption!"

	encrypted := encrypt(plaintext, key)
	fmt.Println("Encrypted:", encrypted)

	decrypted := decrypt(encrypted, key)
	fmt.Println("Decrypted:", decrypted)
}
