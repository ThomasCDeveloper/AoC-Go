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

func applyWindow(ran []int, start int, end int) []int {
	output := []int{}

	for i := 0; i < len(ran); i = i + 2 {
		s := ran[i]
		e := ran[i+1]

		// CHANGE HERE TWEAK
		if s >= end || e < start || (s >= start && e < end) {
			output = append(output, s, e)
		}
		if s < start && e >= start {
			output = append(output, s, start-1, start, e)
		}
		if s < end && e >= end {
			output = append(output, s, end-1, end, e)
		}
	}

	return output
}

func processRange(ran []int, ma ConvertMap) []int {

	sections := ran
	for i := 0; i < len(ran); i = i + 2 {
		for _, mapRange := range ma.ranges {
			mapRangeStart := mapRange.from
			mapRangeEnd := mapRange.from + mapRange.n

			sections = applyWindow(sections, mapRangeStart, mapRangeEnd)

		}
	}

	results := []int{}

	for _, val := range sections {
		results = append(results, processSeed(val, ma))
	}

	return results
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

		if strings.Contains("abcdefghijklmnopqrstuvwxyz", string(line[0])) {
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

	maps := []ConvertMap{}
	newMap := ConvertMap{}
	for _, line := range data[2:] {
		if line == "" {
			continue
		}

		if strings.Contains("abcdefghijklmnopqrstuvwxyz", string(line[0])) {
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

	for i := 0; i < len(maps); i++ {
		fmt.Println(seeds)
		seeds = processRange(seeds, maps[i])
	}

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

func main() {
	data := GetInput("test.txt")

	// PART 1
	//fmt.Println("Part 1:")
	//fmt.Println(SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:")
	fmt.Println(SolvePart2(data))
}
