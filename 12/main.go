package main

import (
	"log"
	"os"
	"slices"
	"strings"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	log.SetFlags(0)
	filename := os.Args[1]
	lines, err := aoc.ReadLinesFromFile(filename)
	if err != nil {
		panic(err.Error())
	}
	if err := Task1(lines); err != nil {
		panic(err.Error())
	}
}

func Task1(lines []string) error {
	total := 0
	for _, line := range lines {
		rawSplit := strings.Split(line, " ")
		groups, err := aoc.ReadIntsFromString(strings.ReplaceAll(rawSplit[1], ",", " "))
		if err != nil {
			return err
		}
		raw := rawSplit[0]
		unknowns := findUnknown(raw)
		arr := validCombinations([]rune(raw), unknowns, groups)
		log.Printf("%d %s %v", arr, raw, groups)
		total += arr
	}
	log.Println(total)
	return nil
}

func findUnknown(s string) []int {
	unknowns := make([]int, 0)
	for i, c := range s {
		if c == '?' {
			unknowns = append(unknowns, i)
		}
	}
	return unknowns
}

func validCombinations(raw []rune, unknowns []int, groups []int) int {
	replacements := make([]rune, 0)
	for i, u := range unknowns {
		raw[u] = replacements[i]
	}

	if valid(string(raw), groups) {
		return 1
	}
	return 0
}

func starting(l int, groups []int) [][]rune {
	curr := make([]rune, l)
	g := 0
	c := 0
	for i := 0; i < l; i++ {
		if g == len(groups) {
			curr[i] = '.'
			continue
		}
		if c < groups[g] {
			curr[i] = '#'
			c++
			continue
		}
		if c == groups[g] {
			curr[i] = '.'
			g++
			c = 0
			continue
		}
	}
	return [][]rune{curr}
}

func valid(raw string, groups []int) bool {
	damaged := make([]int, 0)
	count := 0
	for _, char := range raw {
		if char == '#' {
			count++
			continue
		}
		if count > 0 {
			damaged = append(damaged, count)
			count = 0
		}
	}
	damaged = append(damaged, count)
	return slices.Equal(damaged, groups)
}

func Task2(lines []string) error {
	return nil
}
