package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lukafilipdev/aoc/utils"
)

func main() {
	content, err := utils.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	count := 0
	crossCount := 0

	for i, line := range lines {
		for j, rune := range line {

			if rune == 'X' || rune == 'S' {
				count += checkHorizontal(line, j)
				count += checkVertical(lines, i, j)
				count += checkRightDiagonal(lines, i, j)
				count += checkLeftDiagonal(lines, i, j)
			}

			if i == 0 || j == 0 || i == len(lines) - 1 || j == len(line) - 1 {
				continue
			}

			if rune == 'A' {
				crossCount += checkCross(lines, i, j)
			}

		}
	}

	fmt.Println(count)
	fmt.Println(crossCount)
}

func checkHorizontal(line string, j int) int {

	if j + 3 >= len(line) {
		return 0
	}

	if line[j] == 'X' && line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S' {
		return 1
	}

	if line[j] == 'S' && line[j+1] == 'A' && line[j+2] == 'M' && line[j+3] == 'X' {
		return 1
	}

	return 0
}

func checkVertical(lines []string, i int, j int) int {

	if i + 3 >= len(lines) {
		return 0
	}

	if lines[i][j] == 'X' && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
		return 1
	}

	if lines[i][j] == 'S' && lines[i+1][j] == 'A' && lines[i+2][j] == 'M' && lines[i+3][j] == 'X' {
		return 1
	}

	return 0
}

func checkRightDiagonal(lines []string, i int, j int) int {

	if i + 3 >= len(lines) || j + 3 >= len(lines[i]) {
		return 0
	}

	if lines[i][j] == 'X' && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
		return 1
	}

	if lines[i][j] == 'S' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M' && lines[i+3][j+3] == 'X' {
		return 1
	}

	return 0
}

func checkLeftDiagonal(lines []string, i int, j int) int {

	if i + 3 >= len(lines) || j - 3 < 0 {
		return 0
	}

	if lines[i][j] == 'X' && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
		return 1
	}

	if lines[i][j] == 'S' && lines[i+1][j-1] == 'A' && lines[i+2][j-2] == 'M' && lines[i+3][j-3] == 'X' {
		return 1
	}

	return 0
}

func checkCross(lines []string, i int, j int) int {
	if checkLeftCross(lines, i, j) && checkRightCross(lines, i, j) {
		return 1
	}

	return 0
}

func checkLeftCross(lines []string, i int, j int) bool {
	if lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S' {
		return true
	}

	if lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M' {
		return true
	}

	return false
}

func checkRightCross(lines []string, i int, j int) bool {
	if lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' {
		return true
	}

	if lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' {
		return true
	}

	return false
}
