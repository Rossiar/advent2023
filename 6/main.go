package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	filename := os.Args[1]
	lines, err := aoc.ReadLinesFromFile(filename)
	if err != nil {
		panic(err.Error())
	}
	if err := task(lines); err != nil {
		panic(err.Error())
	}
}

func task(lines []string) error {
	times, err := aoc.ReadIntsFromString(strings.TrimPrefix(lines[0], "Time:"))
	if err != nil {
		return err
	}
	dists, err := aoc.ReadIntsFromString(strings.TrimPrefix(lines[1], "Distance:"))
	if err != nil {
		return err
	}
	if len(times) != len(dists) {
		return fmt.Errorf("mismatched times and distances")
	}
	total := 1
	for i := 0; i < len(times); i++ {
		waysToWin := 0
		raceLength := times[i]
		for timeHeld := 1; timeHeld <= raceLength; timeHeld++ {
			timeToRace := raceLength - timeHeld
			dist := timeHeld * timeToRace
			if dist > dists[i] {
				waysToWin++
			}
		}
		total *= waysToWin
	}
	log.Printf("total: %d", total)
	return nil
}
