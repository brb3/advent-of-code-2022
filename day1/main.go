package main

import (
	"advent_utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	lines := advent_utils.LoadData("input.txt")
	calorieCounts := buildDescSortedCalorieCounts(lines)
	println(calorieCounts[0])
	println(calorieCounts[0] + calorieCounts[1] + calorieCounts[2])
}

func buildDescSortedCalorieCounts(lines []string) []int64 {
	var caloriesCount []int64
	var currentTotal int64 = 0

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			caloriesCount = append(caloriesCount, currentTotal)
			currentTotal = 0
			continue
		}

		line, err := strconv.ParseInt(lines[i], 10, 0)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse line %d", i+1))
		}
		currentTotal += line
	}

	sort.Slice(caloriesCount, func(i int, j int) bool {
		return caloriesCount[i] > caloriesCount[j]
	})
	return caloriesCount
}
