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
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	rows := len(lines)
	cols := len(lines[0])

	grid := make([][]rune, rows)
	var startRow, startCol int

	startDir := 0

	for i, row := range lines {
		grid[i] = []rune(lines[i])
		for j := range row {
			ch := grid[i][j]
			if ch == '^' {
				startRow, startCol = i, j
				grid[i][j] = '.'
			}
		}
	}

	visited := simulate(grid, startRow, startCol, startDir)
	fmt.Println(len(visited))

	loopCount := 0

	for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {

            if (r == startRow && c == startCol) {
                continue
            }

            if grid[r][c] == '.' {
                grid[r][c] = '#'
                if causesLoop(grid, startRow, startCol, startDir) {
                    loopCount++
                }
                grid[r][c] = '.'
            }
        }
    }

	fmt.Println(loopCount)
}

var directions = [4][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func simulate(grid [][]rune, startRow, startCol, startDir int) map[[2]int]bool {
	rows := len(grid)
	cols := len(grid[0])

	row, col := startRow, startCol
	dir := startDir

	visited := make(map[[2]int]bool)
	visited[[2]int{row, col}] = true

	for {
		dx, dy := directions[dir][0], directions[dir][1]
		nr, nc := row+dy, col+dx

		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			break
		}

		if grid[nr][nc] == '#' {
			dir = (dir + 1) % 4
			continue
		}

		row, col = nr, nc
		visited[[2]int{row, col}] = true
	}

	return visited
}

func causesLoop(grid [][]rune, startRow, startCol, startDir int) bool {
	rows := len(grid)
	cols := len(grid[0])

	row, col := startRow, startCol
	dir := startDir

	seen := make(map[[3]int]bool)
	seen[[3]int{row, col, dir}] = true

	for {
		dx, dy := directions[dir][0], directions[dir][1]
		nr, nc := row+dy, col+dx

		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			return false
		}

		if grid[nr][nc] == '#' {
			dir = (dir + 1) % 4
		} else {
			row, col = nr, nc
		}

		state := [3]int{row, col, dir}

		if seen[state] {
			return true
		}

		seen[state] = true
	}
}
