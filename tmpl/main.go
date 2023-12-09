package main

import (
	"os"

	aoc "github.com/rossiar/advent2023"
)

func main() {
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
	return nil
}

func Task2(lines []string) error {
	return nil
}
