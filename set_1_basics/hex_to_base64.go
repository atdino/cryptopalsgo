package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func to_base64(input []byte) string {
	buf := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buf)

	_, err := encoder.Write(input)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	input_hex, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	input_b64 := to_base64(input_hex)
	fmt.Println(input_b64)
}
