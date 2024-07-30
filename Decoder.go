package main
import (
	"fmt"
	"strconv"
	"strings"
)

func Decoder(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var result strings.Builder
	for _, line := range lines {
		line = removeNonPrintables(line)
		decodedLine, err := decodeLine(line)
		if err != nil {
			return "", err
		}
		result.WriteString(decodedLine + "\n")
	}
	return result.String(), nil
}

func decodeLine(line string) (string, error) {
	decoded := ""
	i := 0
	for i < len(line) {
		if line[i] == '[' {
			count, j, char, err := extractCount(line, i)
			if err != nil {
				return "", err
			}
			for k := 0; k < count; k++ {
				decoded += char
			}
			i = j + 1
		} else {
			decoded += string(line[i])
			i++
		}
	}
	return decoded, nil
}

func extractCount(line string, index int) (int, int, string, error) {
	count := ""
	i := index + 1
	for i < len(line) && line[i] >= '0' && line[i] <= '9' {
		count += string(line[i])
		i++
	}
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return 0, 0, "", fmt.Errorf("first argument is not a number")
	}

	if i >= len(line) || line[i] != ' ' {
		return 0, 0, "", fmt.Errorf("No space between arguments")
	}
	i++

	var argBuffer strings.Builder
	openBrackets := 0
	for i < len(line) && line[i] != ']' {
		if line[i] == '[' {
			openBrackets++
		} else if line[i] == ']' {
			if openBrackets == 0 {
				return 0, 0, "", fmt.Errorf("unbalanced brackets within argument")
			}
			openBrackets--
		}
		argBuffer.WriteByte(line[i])
		i++
	}
	if openBrackets != 0 {
		return 0, 0, "", fmt.Errorf("unbalanced brackets within argument")
	}

	arg := argBuffer.String()
	if arg == "" {
		return 0, 0, "", fmt.Errorf("no second argument")
	}

	if i >= len(line) || line[i] != ']' {
		return 0, 0, "", fmt.Errorf("missing closing bracket for argument")
	}

	return countInt, i, argBuffer.String(), nil
}
