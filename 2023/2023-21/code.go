package main

// CMD: go run *.go

import (
	"fmt"
	"slices"
)

type co struct {
	r, c int
}

var tiles = map[co]bool{}

func printSituation(list []co) {
	fmt.Println()
	for r := range data {
		line := ""
		for c := range data[0] {
			if slices.Contains(list, co{r, c}) {
				line += "O"
			} else {
				if tiles[co{r, c}] {
					line += " "
				} else {
					line += "#"
				}
			}
		}
		fmt.Println(line)
	}
}

func deepCopy(list []co) []co {
	out := []co{}
	out = append(out, list...)
	return out
}

func getTile(rc co, loop bool) bool {
	if loop {
		for rc.c < 0 {
			rc.c += len(data)
		}
		for rc.r < 0 {
			rc.r += len(data[0])
		}
		return tiles[co{rc.r % len(data), rc.c % len(data[0])}]
	}
	if rc.c < 0 || rc.r < 0 || rc.r >= len(data) || rc.c >= len(data[0]) {
		return false
	}
	return tiles[rc]
}

func SolvePart1(data []string) int {
	startCo := co{0, 0}
	for r, line := range data {
		for c, char := range line {
			rcco := co{r, c}
			if char == 'S' {
				startCo = co{r, c}
			}
			if char == '#' {
				tiles[rcco] = false
			} else {
				tiles[rcco] = true
			}
		}
	}

	visitedTiles := []co{startCo}

	for i := 0; i < 6; i++ {
		newVisitedTiles := []co{}
		for _, visitedTile := range visitedTiles {
			for _, dir := range []co{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
				rc := co{visitedTile.r + dir.r, visitedTile.c + dir.c}
				if getTile(rc, false) {
					if !slices.Contains(newVisitedTiles, rc) {
						newVisitedTiles = append(newVisitedTiles, rc)
					}
				}
			}
		}
		visitedTiles = newVisitedTiles
	}

	return len(visitedTiles)
}

func SolvePart2(data []string) int {
	startCo := co{0, 0}
	for r, line := range data {
		for c, char := range line {
			rcco := co{r, c}
			if char == 'S' {
				startCo = co{r, c}
			}
			if char == '#' {
				tiles[rcco] = false
			} else {
				tiles[rcco] = true
			}
		}
	}

	visitedTilesBefore := []co{}
	visitedTiles := []co{startCo}
	newVisitedTiles := []co{}
	sum := 0

	nbPas := 0

	for i := 0; i < 65+131; i++ {
		nbPas++

		newVisitedTiles = []co{}

		for _, visitedTile := range visitedTiles {
			for _, dir := range []co{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
				rc := co{visitedTile.r + dir.r, visitedTile.c + dir.c}
				if getTile(rc, true) {
					if !slices.Contains(visitedTilesBefore, rc) {
						if !slices.Contains(newVisitedTiles, rc) {
							newVisitedTiles = append(newVisitedTiles, rc)
						}
					}
				}
			}
		}

		if i%2 == 1 {
			sum += len(newVisitedTiles)
		}
		if i%131 == 0 {
			fmt.Println(i, sum)
		}

		visitedTilesBefore = deepCopy(visitedTiles)
		visitedTiles = deepCopy(newVisitedTiles)
	}

	for i := 0; i < 202299; i++ {
		nbPas += 131
		sum += 948 + (i+1)*472
	}

	return sum
}

var data = []string{}

func main() {
	data = GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
