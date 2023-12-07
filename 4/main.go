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
	total := 0
	for _, line := range lines {
		line := line[strings.Index(line, ":")+2:]
		before, after, _ := strings.Cut(line, " | ")
		winners := strings.Fields(before)
		results := strings.Fields(after)
		score := 0
		for _, winner := range winners {
			if slices.Contains(results, winner) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		fmt.Printf("card with winners %+v and results %+v with score %d\n", winners, results, score)
		total += score
	}
	fmt.Printf("total: %d\n", total)
	return nil
}
