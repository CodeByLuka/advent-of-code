package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/lukafilipdev/aoc/utils"
)

func main() {
	content, err := utils.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	total1 := 0
	total2 := 0
	isMulEnabled := true

	for i := 0; i < len(content)-3; i++ {

		isMulEnabled, i = doInstruction(content, i, isMulEnabled)
		isMulEnabled, i = dontInstruction(content, i, isMulEnabled)

		if content[i] == 'm' && content[i+1] == 'u' && content[i+2] == 'l' && content[i+3] == '(' {
			num1, numLen1, ok := parseNumber(content, i+4)
			if !ok {
				continue
			}

			commaIndex := i + 4 + numLen1
			if commaIndex >= len(content) || content[commaIndex] != ',' {
				continue
			}

			num2, numLen2, ok := parseNumber(content, commaIndex+1)
			if !ok {
				continue
			}

			closingBracketIndex := commaIndex + 1 + numLen2

			if closingBracketIndex >= len(content) || content[closingBracketIndex] != ')' {
				continue
			}

			total1 += num1 * num2

			if isMulEnabled {
				total2 += num1 * num2
			}
		}
	}

	fmt.Println(total1)
	fmt.Println(total2)
}

func parseNumber(data string, startIndex int) (int, int, bool) {
	num := 0
	numLen := 0

	for i := startIndex; i < len(data) && numLen < 3; i++ {
		if !unicode.IsDigit(rune(data[i])) {
			break
		}

		num = num*10 + int(data[i]-'0')
		numLen++
	}

	if numLen == 0 || numLen > 3 {
		numLen = int(numLen)
		return 0, numLen, false
	}

	return num, numLen, true
}

func doInstruction(data string, startIndex int, mulStatus bool) (bool, int) {

	if data[startIndex] == 'd' && data[startIndex+1] == 'o' && data[startIndex+2] == '(' && data[startIndex+3] == ')' {
		return true, startIndex + 4
	}

	return mulStatus, startIndex
}

func dontInstruction(data string, startIndex int, mulStatus bool) (bool, int) {

	if data[startIndex] == 'd' && data[startIndex+1] == 'o' && data[startIndex+2] == 'n' && data[startIndex+3] == '\'' && data[startIndex+4] == 't' && data[startIndex+5] == '(' && data[startIndex+6] == ')' {
		return false, startIndex + 7
	}

	return mulStatus, startIndex
}
