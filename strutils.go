package strutils

import (
	"bytes"
	"unicode"
)

// Swap the case of string parameter

func SwapCase(str string) string {
	
	buf := &bytes.Buffer{}
	for _, r := range str {
		if unicode.IsUpper(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			bun.WriteRune(unicode.ToUpper(r))
		}
	}
	
	return buf.String()
}

// Reverse the string parameter

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)     
}


