package main

// CMD: go run *.go

import (
	"fmt"
	"strings"
)

func load(tiles [][]int) int {
	height := len(tiles)
	sum := 0
	for y := 0; y < len(tiles); y++ {
		for x := 0; x < len(tiles[0]); x++ {
			if tiles[y][x] == 2 {
				sum += height - y
			}
		}
	}
	return sum
}

func trySwap(x int, y int, x2 int, y2 int, tiles [][]int) bool {
	if x < 0 || y < 0 || x2 < 0 || y2 < 0 || x >= len(tiles[0]) || x2 >= len(tiles[0]) || y >= len(tiles) || y2 >= len(tiles) {
		return false
	}
	if tiles[y2][x2] != 0 {
		return false
	}

	tiles[y][x], tiles[y2][x2] = tiles[y2][x2], tiles[y][x]

	return true
}

func MoveUp(tiles [][]int) [][]int {
	for x := 0; x < len(tiles[0]); x++ {
		for y := 0; y < len(tiles); y++ {
			compt := 0
			if tiles[y][x] == 2 {
				for trySwap(x, y-compt, x, y-1-compt, tiles) {
					compt++
				}
			}
		}
	}
	return tiles
}
func MoveDown(tiles [][]int) [][]int {
	for x := 0; x < len(tiles[0]); x++ {
		for y := 0; y < len(tiles); y++ {
			_y := len(tiles) - y - 1
			compt := 0
			if tiles[_y][x] == 2 {
				for trySwap(x, _y+compt, x, _y+1+compt, tiles) {
					compt++
				}
			}
		}
	}
	return tiles
}
func MoveLeft(tiles [][]int) [][]int {
	for x := 0; x < len(tiles[0]); x++ {
		for y := 0; y < len(tiles); y++ {
			compt := 0
			if tiles[y][x] == 2 {
				for trySwap(x-compt, y, x-1-compt, y, tiles) {
					compt++
				}
			}
		}
	}
	return tiles
}
func MoveRight(tiles [][]int) [][]int {
	for x := 0; x < len(tiles[0]); x++ {
		for y := 0; y < len(tiles); y++ {
			_x := len(tiles[0]) - x - 1
			compt := 0
			if tiles[y][_x] == 2 {
				for trySwap(_x+compt, y, _x+1+compt, y, tiles) {
					compt++
				}
			}
		}
	}
	return tiles
}

func MakeSpin(tiles [][]int) [][]int {
	tiles = MoveUp(tiles)
	tiles = MoveLeft(tiles)
	tiles = MoveDown(tiles)
	tiles = MoveRight(tiles)

	return tiles
}

func SolvePart1(data []string) int {
	tiles := [][]int{}
	for _, line := range data {
		tile := strings.Split(line, "")
		tileLine := []int{}
		for j := range tile {
			switch string(tile[j]) {
			case ".":
				tileLine = append(tileLine, 0)
			case "#":
				tileLine = append(tileLine, 1)
			default:
				tileLine = append(tileLine, 2)
			}
		}
		tiles = append(tiles, tileLine)
	}

	MoveUp(tiles)

	return load(MoveUp(tiles))
}

func SolvePart2(data []string) int {
	tiles := [][]int{}
	for _, line := range data {
		tile := strings.Split(line, "")
		tileLine := []int{}
		for j := range tile {
			switch string(tile[j]) {
			case ".":
				tileLine = append(tileLine, 0)
			case "#":
				tileLine = append(tileLine, 1)
			default:
				tileLine = append(tileLine, 2)
			}
		}
		tiles = append(tiles, tileLine)
	}

	for i := 0; i < 1000000000%42+30*42; i++ { // I noticed a pattern of 42
		tiles = MakeSpin(tiles)
	}

	return load(tiles)
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
