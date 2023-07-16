package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
)

func encryptECB(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedText := pad(plainText, block.BlockSize())
	ciphertext := make([]byte, len(paddedText))

	for i := 0; i < len(paddedText); i += block.BlockSize() {
		block.Encrypt(ciphertext[i:i+block.BlockSize()], paddedText[i:i+block.BlockSize()])
	}

	return ciphertext, nil
}

func decryptECB(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(ciphertext))

	for i := 0; i < len(ciphertext); i += block.BlockSize() {
		block.Decrypt(plainText[i:i+block.BlockSize()], ciphertext[i:i+block.BlockSize()])
	}

	return unpad(plainText), nil
}

func encryptOFB(plainText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encryptedIV := make([]byte, len(iv))
	block.Encrypt(encryptedIV, iv)

	ciphertext := make([]byte, len(plainText))

	stream := cipher.NewOFB(block, encryptedIV)
	stream.XORKeyStream(ciphertext, plainText)

	return ciphertext, nil
}

func decryptOFB(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encryptedIV := make([]byte, len(iv))
	block.Encrypt(encryptedIV, iv)

	plainText := make([]byte, len(ciphertext))

	stream := cipher.NewOFB(block, encryptedIV)
	stream.XORKeyStream(plainText, ciphertext)

	return plainText, nil
}

func pad(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	paddedText := append(plainText, bytesRepeating(byte(padding), padding)...)
	return paddedText
}

func unpad(paddedText []byte) []byte {
	padding := int(paddedText[len(paddedText)-1])
	unpaddedText := paddedText[:len(paddedText)-padding]
	return unpaddedText
}

func bytesRepeating(b byte, count int) []byte {
	bytes := make([]byte, count)
	for i := 0; i < count; i++ {
		bytes[i] = b
	}
	return bytes
}

func main() {
	key := []byte("0123456789abcdef")
	iv := []byte("0123456789abcdef")

	var input string
	fmt.Print("Enter the plaintext: ")
	fmt.Scanln(&input)

	plainText := []byte(input)
	//Plaintext
	fmt.Println("Plaintext:", string(plainText))

	ciphertextECB, err := encryptECB(plainText, key)
	if err != nil {
		log.Fatal("ECB Encryption error:", err)
	}

	fmt.Println("Ciphertext (ECB):", hex.EncodeToString(ciphertextECB))

	decryptedTextECB, err := decryptECB(ciphertextECB, key)
	if err != nil {
		log.Fatal("ECB Decryption error:", err)
	}

	fmt.Println("Decrypted text (ECB):", string(decryptedTextECB))

	ciphertextOFB, err := encryptOFB(plainText, key, iv)
	if err != nil {
		log.Fatal("OFB Encryption error:", err)
	}

	fmt.Println("Ciphertext (OFB):", hex.EncodeToString(ciphertextOFB))

	decryptedTextOFB, err := decryptOFB(ciphertextOFB, key, iv)
	if err != nil {
		log.Fatal("OFB Decryption error:", err)
	}

	fmt.Println("Decrypted text (OFB):", string(decryptedTextOFB))
}
