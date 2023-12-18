package main

// CMD: go run *.go

import (
	"fmt"
	"strconv"
	"strings"
)

type co struct {
	r int
	c int
}

func printTiles(tiles [][]bool) {
	for r := 0; r < len(tiles); r++ {
		line := ""
		for c := 0; c < len(tiles[0]); c++ {
			if tiles[r][c] {
				line += "# "
			} else {
				line += "  "
			}
		}
		fmt.Println(line)
	}
}

func SolvePart1(data []string) int {
	listTiles := map[co]bool{}

	currentPos := co{0, 0}
	dirs := map[string]co{"R": {0, 1}, "L": {0, -1}, "U": {-1, 0}, "D": {1, 0}}

	for _, line := range data {
		dir := strings.Split(line, " ")[0]
		n, _ := strconv.Atoi(strings.Split(line, " ")[1])

		for i := 0; i < n; i++ {
			listTiles[currentPos] = true

			currentPos.c += dirs[dir].c
			currentPos.r += dirs[dir].r
		}
	}

	minr, minc, maxr, maxc := currentPos.r, currentPos.c, currentPos.r, currentPos.c
	for key := range listTiles {
		if key.c > maxc {
			maxc = key.c
		}
		if key.c < minc {
			minc = key.c
		}
		if key.r > maxr {
			maxr = key.r
		}
		if key.r < minr {
			minr = key.r
		}
	}

	tiles := [][]bool{}
	for r := 0; r < maxr-minr+1; r++ {
		line := make([]bool, maxc-minc+1)
		tiles = append(tiles, line)
	}

	for key := range listTiles {
		tiles[key.r-minr][key.c-minc] = true
	}
	printTiles(tiles)
	/*
		for r := 0; r < len(tiles); r++ {
			changeWalls := false
			for c := 0; c < len(tiles[0])-1; c++ {
				if tiles[r][c] && !tiles[r][c+1] {
					changeWalls = !changeWalls
				}

				if changeWalls {
					tiles[r][c] = true
				}
			}
		}*/

	sum := 0
	for r := 0; r < len(tiles); r++ {
		for c := 0; c < len(tiles[0]); c++ {
			if tiles[r][c] {
				sum += 1
			}
		}
	}

	return sum
}

func SolvePart2(data []string) int {
	return 0
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
