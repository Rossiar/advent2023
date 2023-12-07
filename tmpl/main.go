package main

import (
	"flag"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	filename := flag.Args()[0]
	lines, err := aoc.ReadLinesFromFile(filename)
	if err != nil {
		panic(err.Error())
	}
	if err := task(lines); err != nil {
		panic(err.Error())
	}
}

func task(lines []string) error {
	return nil
}
