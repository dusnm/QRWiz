package escape

import (
	"encoding/hex"
	"strings"
)

func HexString(input string) string {
	if len(input)%2 != 0 {
		return input
	}

	if _, err := hex.DecodeString(input); err != nil {
		return input
	}

	// The string can be interpreted as a hex number
	// therefore we wrap it in double quotes
	builder := strings.Builder{}
	builder.WriteRune('"')
	builder.WriteString(input)
	builder.WriteRune('"')

	return builder.String()
}

func SpecialCharacters(input string) string {
	builder := strings.Builder{}
	for _, c := range input {
		switch c {
		case '\\', ';', '"', ':':
			builder.WriteRune('\\')
			fallthrough
		default:
			builder.WriteRune(c)
		}
	}

	return builder.String()
}
