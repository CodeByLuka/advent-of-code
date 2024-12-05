package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"

	"github.com/lukafilipdev/aoc/utils"
)

func main() {
	content, err := utils.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	rules := make(map[int][]int)

	total1 := 0
	total2 := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

		line = strings.TrimSpace(line)
		nums := getNums(line)

		switch {
		case strings.Contains(line, "|"):
			processRule(nums, rules)
		case strings.Contains(line, ","):
			processUpdate(nums, rules, &total1, &total2)
		}

	}

	fmt.Println(total1)
	fmt.Println(total2)
}

func processRule(nums []int, rules map[int][]int) {
	rules[nums[0]] = append(rules[nums[0]], nums[1])
}

func processUpdate(nums []int, rules map[int][]int, total1 *int, total2 *int) {
	if isLineCorrect(nums, rules) {
		*total1 += nums[len(nums) / 2]
		return
	}

	correctedNums := correctLine(nums, rules)
	*total2 += correctedNums[len(correctedNums) / 2]
}

func getNums(line string) []int {

	temp := strings.Split(line, string(line[2]))
	nums := make([]int, len(temp))

	for i, numStr := range temp {

		num, err := strconv.Atoi(numStr)

		if err != nil {
			fmt.Println("Error converting:", err)
			os.Exit(1)
		}

		nums[i] = num
	}

	return nums
}

func isLineCorrect(nums []int, rules map[int][]int) bool {
	for i := range nums {
		for _, rule := range rules[nums[i]] {
			for j := 0; j < i; j++ {
				if nums[j] == rule {
					return false
				}
			}
		}
	}

	return true
}

func correctLine(nums []int, rules map[int][]int) []int {
	for i := range nums {
		for _, rule := range rules[nums[i]] {
			for j := 0; j < i; j++ {
				if nums[j] == rule {
					num := nums[i]
					nums = append(nums[:i], nums[i+1:]...)
					nums = append(nums[:j], append([]int{num}, nums[j:]...)...)
					
					return correctLine(nums, rules)
				}
			}
		}
	}

	return nums
}
