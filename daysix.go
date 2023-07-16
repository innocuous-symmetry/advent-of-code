package main

func areUnique(inputs [4]byte) bool {
	switch {
		case inputs[0] == inputs[1]:
		case inputs[0] == inputs[2]:
		case inputs[0] == inputs[3]:
		case inputs[1] == inputs[2]:
		case inputs[1] == inputs[3]:
		case inputs[2] == inputs[3]:
			return false
	}

	return true
}

func daySix(input string) (marker int) {
	var a, b, c, d byte
	marker = 4

	for i := 3; i < len(input); i++ {
		a, b, c, d = input[i-3], input[i-2], input[i-1], input[i]

		if areUnique([4]byte{a, b, c, d}) {
			return
		}

		marker++
	}

	return -1
}
