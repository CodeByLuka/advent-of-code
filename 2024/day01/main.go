package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lukafilipdev/aoc/utils"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	content, err := utils.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	tokens := strings.Fields(content)

	var data []int
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			os.Exit(1)
		}

		data = append(data, num)
	}

	var leftList, rightList []int
	for i, num := range data {
		if i%2 == 0 {
			leftList = append(leftList, num)
		} else {
			rightList = append(rightList, num)
		}
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDiff := 0
	for i := 0; i < len(leftList); i++ {
		diff := abs(leftList[i] - rightList[i])
		totalDiff += diff
	}

	fmt.Println("Total difference:", totalDiff)

	// Part 2
	rightFrequncy := make(map[int]int)
	for _, num := range rightList {
		rightFrequncy[num]++
	}

	totalSimilarity := 0
	for _, num := range leftList {
		totalSimilarity += num * rightFrequncy[num]
	}

	fmt.Println("Total similarity:", totalSimilarity)
}
