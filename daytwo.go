package main

import (
	"fmt"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}

func playValueOf(s string) int {
	fmt.Println("played value: " + s)
	fmt.Println("last char: " + string(s[len(s)-1]))

	switch string(s[len(s)-1]) {
	case "Z":
		return 3
	case "Y":
		return 2
	case "X":
		return 1
	default:
		panic("Invalid input received at func playValueOf")
	}
}

func dayTwo(input string) int {
	// split input by newlines
	var eachPart = strings.Split(input, "\n")

	// trim all remaining whitespace
	for j := range eachPart {
		eachPart[j] = strings.TrimSpace(eachPart[j])
	}

	winningResults := []string{
		"A Y",
		"B Z",
		"C X",
	}

	drawResult := []string{
		"A X",
		"B Y",
		"C Z",
	}

	losingResult := []string{
		"A Z",
		"B X",
		"C Y",
	}

	score := 0

	for i := 0; i <= 3; i++ {
		switch true {
			case contains(winningResults, eachPart[i]):
				score += playValueOf(eachPart[i]) + 6
			case contains(drawResult, eachPart[i]):
				score += playValueOf(eachPart[i]) + 3
			case contains(losingResult, eachPart[i]):
				score += playValueOf(eachPart[i])
			default:
				score += 0
		}
	}

	return score
}

