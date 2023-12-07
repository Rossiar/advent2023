package main

import (
	"fmt"
	"slices"
	"strings"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	lines, err := aoc.ReadLinesFromFile("in.txt")
	if err != nil {
		panic(err.Error())
	}
	if err := task(lines); err != nil {
		panic(err.Error())
	}
}

func task(lines []string) error {
	cards := make([]int, len(lines))
	for i, line := range lines {
		line := line[strings.Index(line, ":")+2:]
		before, after, _ := strings.Cut(line, " | ")
		winners := strings.Fields(before)
		results := strings.Fields(after)
		score := 0
		for _, winner := range winners {
			if slices.Contains(results, winner) {
				score++
			}
		}
		cards[i] = score
	}
	total := 0
	for card := range cards {
		total += win(cards, card)
	}
	fmt.Printf("total: %d\n", total)
	return nil
}

func win(cards []int, card int) int {
	total := 1
	score := cards[card]
	//fmt.Printf("scoring card %d, searching next %d cards\n", card+1, score)
	for i := card + 1; i <= card+score; i++ {
		total += win(cards, i)
	}
	return total
}
