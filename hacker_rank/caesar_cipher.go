package hacker_rank

import (
	"unicode"
)

// given a string, and how many characters you want to rotate it, rotate each character in the string
func CaesarCipher(s string, k int) string {

	runes := []rune(s)
	rv := ""
	for i := 0; i < len(s); i++ {
		if runes[i] < 65 || (runes[i] > 90 && runes[i] < 97) || runes[i] > 122 {
			rv += string(runes[i])
			continue
		}
		isCaps := false
		if unicode.IsUpper(runes[i]) {
			isCaps = true
			runes[i] = unicode.ToLower(runes[i])
		}
		if k > 26 {
			k %= 26
		}
		runes[i] += int32(k)
		if runes[i] > 122 {
			runes[i] -= 26
		}
		if isCaps {
			runes[i] = unicode.ToUpper(runes[i])
		}
		rv += string(runes[i])
	}
	return rv
}
