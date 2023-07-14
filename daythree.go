package main

import (
	"fmt"
	"math"
	"strings"
)

func printRunes() {
	fmt.Println("lowercase")
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, v := range alphabet {
		fmt.Println(string(v), v)
	}

	fmt.Println("")
	fmt.Println("uppercase")
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, v := range alphabet {
		fmt.Println(string(v), v)
	}
}

func compare(f string, s string) (result rune) {
	minLength := math.Min(float64(len(f)), float64(len(s)))

	for i := 0; i < int(minLength); i++ {
		// check each character of each string
		if strings.Contains(s, string(f[i])) {
			result = rune(f[i])
			maxValue := math.Max(float64(result), float64(f[i]))

			// assign the greatest value (between the current stored value
			// and the newly found value) to our result
			result = rune(maxValue)
		} else if strings.Contains(f, string(s[i])) {
			result = rune(s[i])
			maxValue := math.Max(float64(result), float64(s[i]))
			result = rune(maxValue)
		}
	}

	return
}

func unique(e string) (result string, duplicated rune) {
	for _, v := range e {
		if strings.Contains(result, string(v)) {
			duplicated = v
			continue
		}

		result += string(v)
	}

	return
}

func runeToScore(input rune) int32 {
	// filter out non-alphabetical runes
	isBetweenRanges := input > 90 && input < 97
	isAboveRange := input > 122
	isBelowRange := input < 65

	// for these, return zero (will not affect total score of solution)
	if isBetweenRanges || isAboveRange || isBelowRange {
		fmt.Println(input)
		return 0
	}

	if input <= 90 {
		// is uppercase
		return int32(input - 38)
	} else {
		// is lowercase
		return int32(input - 96)
	}
}

func dayThree(input string) (solution int32) {
	// split input by newlines
	var eachPart = strings.Split(input, "\n")

	// trim all remaining whitespace
	for j := range eachPart {
		eachPart[j] = strings.TrimSpace(eachPart[j])
	}

	for i := 0; i < len(eachPart); i++ {
		part := eachPart[i]

		if part == "" {
			continue
		}

		// split string into two halves
		front, back := part[0:len(part)/2], part[len(part)/2:]

		fmt.Println(part)
		fmt.Println(front, back)

		// check each character against each other string
		result := compare(front, back)
		score := runeToScore(result)
		fmt.Println(string(result), result, score)

		// add each value to total score
		solution += score
		fmt.Println("")
	}

	return
}
