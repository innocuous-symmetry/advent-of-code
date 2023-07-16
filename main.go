package main

import (
	"fmt"
)

func main() {
	fmt.Println("ADVENT OF CODE 2022 SOLUTIONS -- GOLANG")

	/** DAY ONE */
	fmt.Println("\nDAY ONE")

	var dayOneResult = dayOne(
	`1000
	2000
	3000

	4000

	5000
	6000

	7000
	8000
	9000

	10000`)

	fmt.Println(dayOneResult)

	dayOneResult = dayOne(
	`245
	256
	939

	-12039


	4949

	2
	2
	2`)

	fmt.Println(dayOneResult)

	/** DAY TWO */
	fmt.Println("\nDAY TWO")

	var dayTwoResult = dayTwo(
		`
		A Y
		B X
		C Z
		`,
	)
	fmt.Println(dayTwoResult)

	/** DAY THREE */
	fmt.Println("\nDAY THREE")

	var dayThreeResult = dayThree(
	`
	vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw
	`)

	fmt.Println(dayThreeResult)


	/** DAY FOUR */
	fmt.Println("\nDAY FOUR")

	dayFourInput := [][2][2]int{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}

	var dayFourResult = dayFour(dayFourInput)
	fmt.Println(dayFourResult)


	/** DAY FIVE */
	fmt.Println("\nDAY FIVE")

	dayFiveInput := [4][3]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}

	var dayFiveResult = dayFive(dayFiveInput)
	for _, r := range dayFiveResult {
		fmt.Printf("%c ", r)
	}


	/** DAY SIX */
	fmt.Println("\nDAY SIX")
	fmt.Println(daySix("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	fmt.Println(daySix("nppdvjthqldpwncqszvftbrmjlhg"))
	fmt.Println(daySix("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	fmt.Println(daySix("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
