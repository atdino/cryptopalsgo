package main

import (
	"encoding/hex"
	"fmt"
)

func xor(input_1 []byte, input_2 []byte) { //inplace xor
	for i := range input_1 {
		input_1[i] = input_1[i] ^ input_2[i]
	}
}

func main() {
	input := "1c0111001f010100061a024b53535009181c"
	xor_against := "686974207468652062756c6c277320657965"

	input_hex, _ := hex.DecodeString(input)
	xor_against_hex, _ := hex.DecodeString(xor_against)

	xor(input_hex, xor_against_hex)
	fmt.Printf("as str: %s\n", string(input_hex))
	fmt.Printf("as hex: %x\n", input_hex)
}
