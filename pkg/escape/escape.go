package escape

import (
	"encoding/hex"
	"fmt"
)

func HexString(input string) string {
	if len(input)%2 != 0 {
		return input
	}

	if _, err := hex.DecodeString(input); err == nil {
		// The string can be interpreted as a hex number
		// therefore we wrap it in double quotes
		input = fmt.Sprintf("\"%s\"", input)
	}

	return input
}

func SpecialCharacters(input string) string {
	runes := make([]rune, 0, len(input))
	for _, c := range input {
		if c == '\\' ||
			c == ';' ||
			c == ',' ||
			c == '"' ||
			c == ':' {
			runes = append(runes, '\\')
		}

		runes = append(runes, c)
	}

	return string(runes)
}
