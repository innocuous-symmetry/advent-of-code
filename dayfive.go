package main

import "fmt"

// define stack
type Stack[T any] struct {
	elements []T
}

// add methods
func (s *Stack[T]) Push(elements ...T) {
	s.elements = append(s.elements, elements...)
}

func (s *Stack[T]) Pop() T {
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *Stack[T]) Peek() T {
	return s.elements[len(s.elements)-1]
}

func dayFive(instructions [4][3]int) [3]rune {
	// initialize variables
	stackOne := Stack[rune]{}
	stackOne.Push('Z', 'N')

	stackTwo := Stack[rune]{}
	stackTwo.Push('M', 'C', 'D')

	stackThree := Stack[rune]{}
	stackThree.Push('P')

	stacks := [3]*Stack[rune]{&stackOne, &stackTwo, &stackThree}

	// initial state
	fmt.Println(*stacks[0], *stacks[1], *stacks[2])

	// perform each instruction
	for _, instruction := range instructions {
		sourceStack := stacks[instruction[1]-1]
		targetStack := stacks[instruction[2]-1]
		numToMove := instruction[0]

		for i := 0; i < numToMove; i++ {
			// remove the thing before attempting to move it
			thingToMove := sourceStack.Pop()
			targetStack.Push(thingToMove)
		}

		// log after each step
		fmt.Println(*stacks[0], *stacks[1], *stacks[2])
	}

	return [3]rune{stackOne.Peek(), stackTwo.Peek(), stackThree.Peek()}
}
