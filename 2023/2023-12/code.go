package main

// CMD: go run *.go

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type sequence struct {
	springs string
	cmd     string
}

var cache = make(map[sequence]int)

func GetCombinaisons(cmd []uint8, springs string) int {
	if len(cmd) == 0 && len(springs) == 0 {
		return 1
	}
	if len(springs) == 0 {
		return 0
	}

	value, test := cache[sequence{springs, string(cmd)}]
	if test { // si la sequence est déjà dans le cache
		return value
	}

	if springs[0] == '.' {
		val := GetCombinaisons(cmd, springs[1:])
		cache[sequence{springs, string(cmd)}] = val
		return val
	}

	sum := 0
	for _, n := range cmd {
		sum += int(n)
	}
	if len(springs) < sum+len(cmd)-1 {
		cache[sequence{springs, string(cmd)}] = 0
		return 0
	}

	if springs[0] == '?' {
		res := GetCombinaisons(cmd, springs[1:]) + GetCombinaisons(cmd, "#"+springs[1:])
		cache[sequence{springs, string(cmd)}] = res
		return res
	}

	if springs[0] == '#' {
		if len(cmd) == 0 {
			cache[sequence{springs, string(cmd)}] = 0
			return 0
		}

		n := cmd[0]
		indexDot := strings.Index(springs, ".")
		if indexDot == -1 {
			indexDot = len(springs)
		}
		if indexDot < int(n) {
			cache[sequence{springs, string(cmd)}] = 0
			return 0
		}

		remaining := springs[n:]
		if len(remaining) == 0 {
			res := GetCombinaisons(cmd[1:], remaining)
			cache[sequence{springs, string(cmd)}] = res
			return res
		}

		if remaining[0] == '#' {
			cache[sequence{springs, string(cmd)}] = 0
			return 0
		}

		res := GetCombinaisons(cmd[1:], remaining[1:])
		cache[sequence{springs, string(cmd)}] = res
		return res
	}

	return sum
}

func SolvePart1(data []string) int {
	cmds := [][]uint8{}
	springss := []string{}
	for _, line := range data {
		re := regexp.MustCompile("[0-9]+")

		cmd := []uint8{}
		for _, stringcmd := range re.FindAllString(line, -1) {
			intcmd, _ := strconv.Atoi(stringcmd)
			int8cmd := uint8(intcmd)
			cmd = append(cmd, int8cmd)
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
	cmds := [][]uint8{}
	springss := []string{}
	for _, line := range data {
		re := regexp.MustCompile("[0-9]+")

		cmd := []uint8{}
		for _, stringcmd := range re.FindAllString(line, -1) {
			intcmd, _ := strconv.Atoi(stringcmd)
			int8cmd := uint8(intcmd)
			cmd = append(cmd, int8cmd)
		}

		cmds = append(cmds, cmd)
		springss = append(springss, strings.Split(line, " ")[0])
	}

	sum := 0
	for i := range springss {
		cmd := []uint8{}
		springs := ""
		for j := 0; j < 5; j++ {
			cmd = append(cmd, cmds[i]...)
			springs += "?" + springss[i]
		}
		sum += GetCombinaisons(cmd, springs[1:])
	}

	return sum
}

func main() {
	data := GetInput("test.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
