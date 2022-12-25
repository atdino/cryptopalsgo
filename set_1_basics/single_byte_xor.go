package main

import (
	"encoding/hex"
	"fmt"
    "strings"
)

func xor_against_char(input_1 []byte, char rune) []byte { //inplace xor
    result := append([]byte(nil), input_1...)
	for i := range input_1 {
		result[i] = input_1[i] ^ byte(char)
	}

    return result
}


func reverse(s string) string {
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

        rns[i], rns[j] = rns[j], rns[i]
    }

    return string(rns)
}


func get_score(b []byte) int {
    score_str := "ETAOIN SHRDLU"
    scoring_map := reverse(score_str)
    scoring_map_lowered := reverse(strings.ToLower("ETAOIN SHRDLU"))
    score := 0
    for _, c := range b{
        idx := strings.Index(scoring_map, string(c))
        idx_l := strings.Index(scoring_map_lowered, string(c))
        //fmt.Println(idx)
        score += idx
        score += idx_l
    }
    return score
}

func get_max_score(m map[int]string) int {
    max_key := -1
    for k := range m {
        if (k > max_key) {
            max_key = k
        }
    }
    return max_key
}

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	input_hex, _ := hex.DecodeString(input)

    possible_chars := "abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    str_scores := make(map[int]string)
    // scoring here...
    for _, char := range possible_chars {
        xored_input := xor_against_char(input_hex, char)
        //xored_input_str := string(xored_input)
        score := get_score(xored_input)
        str_scores[score] = string(xored_input)
        //fmt.Printf("[%c] [%d] %s\n", char, score, string(xored_input))
        //fmt.Println("================================")
    }
    max_score := get_max_score(str_scores)
    fmt.Println(str_scores[max_score])

}
