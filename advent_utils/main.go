package advent_utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func LoadData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open input file!")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var content []string
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	return content
}

func ConvertStringArrayToIntArray(array []string) []int {
	var intLines []int
	for i := 0; i < len(array); i++ {
		intValue, err := strconv.ParseInt(array[i], 10, 0)
		if err != nil {
			fmt.Println("Unable to parse string as int")
			os.Exit(1)
		}
		intLines = append(intLines, int(intValue))
	}
	return intLines
}
