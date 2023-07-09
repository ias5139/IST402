package main

import (
	"fmt"
	"strings"
)

func caesarDecrypt(ciphertext string, shift int) string {
	var result strings.Builder
	shift = shift % 26 // Ensure the shift value is within the range of 0-25

	for _, char := range ciphertext {
		if char >= 'a' && char <= 'z' {
			char = 'a' + ((char - 'a' - rune(shift) + 26) % 26)
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + ((char - 'A' - rune(shift) + 26) % 26)
		}
		result.WriteRune(char)
	}

	return result.String()
}

func bruteForceDecrypt(ciphertext string) {
	for shift := 0; shift < 26; shift++ {
		decrypted := caesarDecrypt(ciphertext, shift)
		fmt.Printf("Shift %d: %s\n", shift, decrypted)
	}
}
func main() {
	encryptedSentence := "Ugew gnwj zwjw Oslkgf"
	bruteForceDecrypt(encryptedSentence)
}
