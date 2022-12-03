package main

import (
	"advent_utils"
	"strings"
)

// I don't really like the switch statements (and nested!) but day 3 is already out so moving on
func main() {
	lines := advent_utils.LoadData("input.txt")
	println(dayOne(lines))
	println(dayTwo(lines))
}

func dayOne(rounds []string) int {
	score := 0

	for i := 0; i < len(rounds); i++ {
		fields := strings.Fields(rounds[i])
		theirs := fields[0]
		mine := fields[1]

		switch mine {
		case "X":
			if theirs == "A" { // Draw
				score += 3
			} else if theirs == "C" { // Win
				score += 6
			}
			score += 1
		case "Y":
			if theirs == "B" { // Draw
				score += 3
			} else if theirs == "A" { // Win
				score += 6
			}
			score += 2
		case "Z":
			if theirs == "C" { // Draw
				score += 3
			} else if theirs == "B" { // Win
				score += 6
			}
			score += 3
		}
	}

	return score
}

func dayTwo(rounds []string) int {
	score := 0

	for i := 0; i < len(rounds); i++ {
		fields := strings.Fields(rounds[i])
		theirs := fields[0]
		strat := fields[1]

		switch strat {
		case "X": // Lose
			switch theirs {
			case "A":
				score += 3 // Chose scissors
			case "B":
				score += 1 // Chose rock
			case "C":
				score += 2 // Chose paper
			}
		case "Y": // Draw, choose same as opponent
			score += 3
			switch theirs {
			case "A":
				score += 1
			case "B":
				score += 2
			case "C":
				score += 3
			}
		case "Z": // Win
			score += 6
			switch theirs {
			case "A":
				score += 2 // Chose paper
			case "B":
				score += 3 // Chose scissors
			case "C":
				score += 1 // Chose rock
			}
		}
	}

	return score
}
