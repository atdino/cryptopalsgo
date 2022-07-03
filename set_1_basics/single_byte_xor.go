package main

import (
	"encoding/hex"
	"fmt"
)

func xor_against_char(input_1 []byte, char rune) []byte { //inplace xor
    result := append([]byte(nil), input_1...)
	for i := range input_1 {
		result[i] = input_1[i] ^ byte(char)
	}

    return result
}

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	input_hex, _ := hex.DecodeString(input)

    possible_chars := "abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    for _, char := range possible_chars {
        xored_input := xor_against_char(input_hex, char)
        fmt.Printf("[%c] %s\n", char, string(xored_input))
    }
}
