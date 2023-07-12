package main

import (
	"math"
	"strconv"
	"strings"
)

func dayOne(input string) (maxCalories int) {
	// split input by newlines
	var eachPart = strings.Split(input, "\n")

	// trim all remaining whitespace
	for i := range eachPart {
		eachPart[i] = strings.TrimSpace(eachPart[i])
	}

	// assign a max value and a current value
	maxCalories = 0
	temp := 0

	for i := 0; i < len(eachPart); i++ {
		// this indicates we have reached the end of each elf's snack stash
		if eachPart[i] == "" {
			// use the max between curent temp value and the current max value found
			maxCalories = int(math.Max(float64(maxCalories), float64(temp)))

			// reassign temp to zero and start again
			temp = 0
			continue
		}

		// string to int conversion
		partToInt, err := strconv.Atoi(eachPart[i])

		if err != nil {
			panic(err)
		}

		// add the result to our current working total
		temp += partToInt
	}

	return
}
