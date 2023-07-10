def caesar_decrypt(ciphertext, shift):
    result = ''
    shift = shift % 26  # Ensure the shift value is within the range of 0-25

    for char in ciphertext:
        if 'a' <= char <= 'z':
            char = chr((ord(char) - ord('a') - shift) % 26 + ord('a'))
        elif 'A' <= char <= 'Z':
            char = chr((ord(char) - ord('A') - shift) % 26 + ord('A'))
        result += char

    return result


def brute_force_decrypt(ciphertext):
    for shift in range(26):
        decrypted = caesar_decrypt(ciphertext, shift)
        print(f"Shift {shift}: {decrypted}")


def main():
    encrypted_sentence = "Jxu f qilmtb bw jxu ymjwjx"
    brute_force_decrypt(encrypted_sentence)


if __name__ == '__main__':
    main()