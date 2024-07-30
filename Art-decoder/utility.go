package main

import (
	"strings"
	"unicode"
)

func removeNonPrintables(input string) string {
	var sanitized strings.Builder
	for _, r := range input {
		if unicode.IsPrint(r) {
			sanitized.WriteRune(r)
		}
	}
	return sanitized.String()
}
