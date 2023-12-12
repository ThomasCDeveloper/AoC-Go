package main

// CMD: go run *.go

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func isConform(cmd []int, springs string) bool {
	for strings.Contains(springs, "..") {
		springs = strings.Replace(springs, "..", ".", -1)
	}
	if string(springs[0]) == "." {
		springs = springs[1:]
	}
	if len(springs) > 0 {
		if string(springs[len(springs)-1]) == "." {
			springs = springs[:len(springs)-1]
		}
	}

	springChains := strings.Split(springs, ".")

	if len(springChains) != len(cmd) {
		return false
	}

	for i := range springChains {
		if len(springChains[i]) != cmd[i] {
			return false
		}
	}
	return true
}

func GetCombinaisons(cmd []int, springs string) int {
	splitedSprings := strings.Split(springs, "?")
	numberOfQuestionMarks := len(splitedSprings) - 1

	sum := 0
	for i := 0; i < int(math.Pow(2, float64(numberOfQuestionMarks))); i++ {
		newSprings := ""
		replaceQuestionMarks := strconv.FormatInt(int64(i), 2)
		for len(replaceQuestionMarks) < numberOfQuestionMarks {
			replaceQuestionMarks = "0" + replaceQuestionMarks
		}

		for j := 0; j < len(replaceQuestionMarks); j++ {
			newSprings += splitedSprings[j]
			if string(replaceQuestionMarks[j]) == "1" {
				newSprings += "#"
			} else {
				newSprings += "."
			}
		}
		newSprings += splitedSprings[numberOfQuestionMarks]

		if isConform(cmd, newSprings) {
			sum += 1
		}
	}

	return sum
}

func SolvePart1(data []string) int {
	cmds := [][]int{}
	springss := []string{}
	for _, line := range data {
		re := regexp.MustCompile("[0-9]+")

		cmd := []int{}
		for _, stringcmd := range re.FindAllString(line, -1) {
			intcmd, _ := strconv.Atoi(stringcmd)
			cmd = append(cmd, intcmd)
		}

		cmds = append(cmds, cmd)
		springss = append(springss, strings.Split(line, " ")[0])
	}

	sum := 0
	for i := range springss {
		sum += GetCombinaisons(cmds[i], springss[i])
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
