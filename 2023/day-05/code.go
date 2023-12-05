package main

// CMD: go run *.go

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
	n    int
}

type ConvertMap struct {
	ranges []Range
}

func processSeed(seed int, ma ConvertMap) int {
	for _, ran := range ma.ranges {
		if seed >= ran.from && seed <= ran.from+ran.n {
			return ran.to + seed - ran.from
		}
	}

	return seed
}

func processRange(rang []int, ma ConvertMap) []int {
	output := []int{}
	/*
		for _, ran := range ma.ranges {

		}
	*/
	return output
}

func SolvePart1(data []string) int {
	seeds := []int{}
	for _, val := range strings.Split(data[0], " ")[1:len(strings.Split(data[0], " "))] {
		convert, _ := strconv.Atoi(val)
		seeds = append(seeds, convert)
	}

	maps := []ConvertMap{}
	newMap := ConvertMap{}
	for _, line := range data[2:] {
		if line == "" {
			continue
		}

		if strings.Index("abcdefghijklmnopqrstuvwxyz", string(line[0])) != -1 {
			if len(newMap.ranges) != 0 {
				maps = append(maps, newMap)
				newMap = ConvertMap{}
			}
			continue
		}

		to, _ := strconv.Atoi(strings.Split(line, " ")[0])
		from, _ := strconv.Atoi(strings.Split(line, " ")[1])
		n, _ := strconv.Atoi(strings.Split(line, " ")[2])

		newMap.ranges = append(newMap.ranges, Range{from, to, n})
	}

	maps = append(maps, newMap)

	results := []int{}

	for _, seed := range seeds {
		for _, ma := range maps {
			seed = processSeed(seed, ma)
		}
		results = append(results, seed)
	}

	min := results[0]
	for _, res := range results {
		if res < min {
			min = res
		}
	}

	return min
}

func SolvePart2(data []string) int {
	seeds := []int{}
	splitedLine := strings.Split(data[0], " ")
	for i := 0; i < len(splitedLine)-1; i++ {
		from, _ := strconv.Atoi(splitedLine[i+1])
		n, _ := strconv.Atoi(splitedLine[i+2])
		i++

		seeds = append(seeds, from)
		seeds = append(seeds, from+n)
	}

	fmt.Println(seeds)

	maps := []ConvertMap{}
	newMap := ConvertMap{}
	for _, line := range data[2:] {
		if line == "" {
			continue
		}

		if strings.Index("abcdefghijklmnopqrstuvwxyz", string(line[0])) != -1 {
			if len(newMap.ranges) != 0 {
				maps = append(maps, newMap)
				newMap = ConvertMap{}
			}
			continue
		}

		to, _ := strconv.Atoi(strings.Split(line, " ")[0])
		from, _ := strconv.Atoi(strings.Split(line, " ")[1])
		n, _ := strconv.Atoi(strings.Split(line, " ")[2])

		newMap.ranges = append(newMap.ranges, Range{from, to, n})
	}
	maps = append(maps, newMap)

	/*
		results := []int{}

		for _, seed := range seeds {
			for _, ma := range maps {
				seed = processSeed(seed, ma)
			}
			results = append(results, seed)
		}

		min := results[0]
		for _, res := range results {
			if res < min {
				min = res
			}
		}
	*/

	return 0
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:")
	fmt.Println(SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:")
	fmt.Println(SolvePart2(data))
}
