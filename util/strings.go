package util

import (
	"unicode"
)

func TrimMultiBlank(line string) string {
	res := []rune{}
	for _, v := range line {
		if !unicode.IsSpace(v) {
			res = append(res, v)
		}
	}
	return string(res)
}
