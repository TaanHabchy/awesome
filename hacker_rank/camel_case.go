package hacker_rank

import (
	"unicode"
)

func CamelCase(s string) {
	runes := []rune(s)
	rv := 1
	for i := 0; i < len(s); i++ {
		if unicode.IsUpper(runes[i]) {
			rv += 1
		}
	}
	print(rv)
}
