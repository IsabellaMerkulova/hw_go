package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var b strings.Builder
	var prevChar rune
	var prevIsDigit bool

	for _, char := range input {
		isDigit := unicode.IsDigit(char)
		if isDigit {
			if prevIsDigit || prevChar == 0 {
				return "", ErrInvalidString
			}
			if int(char-'0') == 0 {
				curResultString := b.String()
				b.Reset()
				b.WriteString(curResultString[:len(curResultString)-1])
				continue
			}
			b.WriteString(strings.Repeat(string(prevChar), int(char-'0')-1))
		} else {
			b.WriteRune(char)
		}
		prevChar = char
		prevIsDigit = isDigit
	}
	result := b.String()
	return result, nil
}
