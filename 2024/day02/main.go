package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lukafilipdev/aoc/utils"
)

func isSafe(nums []int) bool {
	increasing, decreasing := true, true

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if diff < 0 {
			increasing = false
		}
		if diff > 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func canBeMadeSafe(nums []int) bool {
	for i := 0; i < len(nums); i++ {

		newNums := append([]int{}, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)

		if isSafe(newNums) {
			return true
		}
	}
	return false
}

func main() {
	content, err := utils.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	data := make([][]int, len(lines))

	for i, line := range lines {
		nums := strings.Fields(line)
		data[i] = make([]int, len(nums))

		for j, num := range nums {
			data[i][j], err = strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				os.Exit(1)
			}
		}
	}

	count1 := 0
	count2 := 0

	for _, nums := range data {
		if isSafe(nums) {
			count1++
		} else if canBeMadeSafe(nums) {
			count2++
		}
	}

	fmt.Println(count1)
	fmt.Println(count1 + count2)
}

