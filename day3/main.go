package main

import (
	"advent_utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := advent_utils.LoadData("input.txt")
	commonItems := findCommonItems(lines)
	println(partOne(commonItems))
}

func partOne(commonItems []string) int64 {
	var sum int64 = 0

	for i := 0; i < len(commonItems); i++ {
		sum += getPriority(commonItems[i])
	}

	return sum
}

func findCommonItems(contents []string) []string {
	commonItems := make([]string, 0)

	for i := 0; i < len(contents); i++ {
		// Split the contents into two compartments
		sizeOfCompartment := len(contents[i]) / 2
		compartmentOne := strings.Split(contents[i][0:sizeOfCompartment], "")
		compartmentTwo := contents[i][sizeOfCompartment:]

		// Find the common item
		for i := 0; i < len(compartmentOne); i++ {
			if strings.Contains(compartmentTwo, compartmentOne[i]) {
				commonItems = append(commonItems, compartmentOne[i])
				break
			}
		}
	}

	return commonItems
}

func getPriority(c string) int64 {
	cv, err := strconv.ParseInt(fmt.Sprintf("%X", c), 16, 0)
	if err != nil {
		panic("Unable to parse character")
	}

	// A-Z is 65-90. Reduce by 38 to bring A to 27
	if cv <= 90 {
		cv -= 38
		return cv
	}

	// a-z is 97-122. Reduce by 96 to bring a to 1
	cv -= 96
	return cv
}
