package main

// CMD: go run *.go

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

type link struct {
	a, b string
}

func GetPath(source Node, sink Node) []Node {

	return []Node{}
}

type Node struct {
	neighbors []string
}

func SolvePart1(data []string) int {
	links := []link{}

	for _, line := range data {
		a := strings.Split(line, ": ")[0]
		for _, b := range strings.Split(strings.Split(line, ": ")[1], " ") {
			links = append(links, link{a, b})
		}
	}

	nodes := map[string]Node{}
	for _, link := range links {
		if _, ok := nodes[link.a]; !ok {
			nodes[link.a] = Node{[]string{}}
		}
		if _, ok := nodes[link.b]; !ok {
			nodes[link.b] = Node{[]string{}}
		}

		if !slices.Contains(nodes[link.a].neighbors, link.b) {
			nodes[link.a] = Node{append(nodes[link.a].neighbors, link.b)}
		}
		if !slices.Contains(nodes[link.b].neighbors, link.a) {
			nodes[link.b] = Node{append(nodes[link.b].neighbors, link.a)}
		}
	}

	//counts := map[link]int{}
	for i := 0; i < 10; i++ {
		// Get sink
		// Get source
		// for node := GetPath {
		//	counts[node]++
		// }
	}

	fmt.Println(rand.Intn(len(links)))
	fmt.Println(nodes)

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
