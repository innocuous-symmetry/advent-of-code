package main

func incrementIfEncapsulated(pair [2][2]int, solution *int) {
	if isEncapsulated(pair[0], pair[1]) {
		*solution++
	}
}

func isEncapsulated(first [2]int, second [2]int) bool {
	return first[0] <= second[0] && first[1] >= second[1] || second[0] <= first[0] && second[1] >= first[1]
}

func dayFour(input [][2][2]int) int {
	solution := 0

	for _, pair := range input {
		incrementIfEncapsulated(pair, &solution)
	}

	return solution
}
