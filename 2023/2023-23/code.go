package main

// CMD: go run *.go

import (
	"fmt"
)

type co struct {
	r, c int
}

var directions = map[rune][]co{
	'.': {{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
	'>': {{1, 0}},
	'<': {{-1, 0}},
	'^': {{0, -1}},
	'V': {{0, 1}},
}

func SolvePart1(data []string) int {

	//startCo := co{1, 1}

	return 0
}

func SolvePart2(data []string) int {
	return 0
}

func main() {
	data := GetInput("test.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
