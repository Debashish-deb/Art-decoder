package main

import (
	"fmt"
	"strings"
)

func Encoder(line string) string {
	lines := strings.Split(line, "\n")

	encodedLines := make([]string, len(lines))
	for i, line := range lines {
		encodedLines[i] = encodeLine(line)
	}

	return strings.Join(encodedLines, "\n")
}

func encodeLine(line string) string {
	line = removeNonPrintables(line)
	encoded := ""
	count := 1
	for i := 0; i < len(line); i++ {
		if i < len(line)-1 && line[i] == line[i+1] {
			count++
		} else {
			if count > 1 {
				encoded += fmt.Sprintf("[%d %c]", count, line[i])
			} else {
				encoded += string(line[i])
			}
			count = 1
		}
	}
	return encoded
}
