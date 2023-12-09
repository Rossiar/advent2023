package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	filename := os.Args[1]
	lines, err := aoc.ReadLinesFromFile(filename)
	if err != nil {
		panic(err.Error())
	}
	if err := task2(lines); err != nil {
		panic(err.Error())
	}
}

func task1(lines []string) error {
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

func task2(lines []string) error {
	clean := func(line, prefix string) (int, error) {
		trimmed := strings.TrimPrefix(line, prefix)
		cleaned := strings.ReplaceAll(trimmed, " ", "")
		return strconv.Atoi(cleaned)
	}
	raceLength, err := clean(lines[0], "Time:")
	if err != nil {
		return err
	}
	recordDist, err := clean(lines[1], "Distance:")
	if err != nil {
		return err
	}
	waysToWin := 0
	lastDist := -1
	i := 0
	for timeHeld := 1; timeHeld <= raceLength; timeHeld++ {
		timeToRace := raceLength - timeHeld
		dist := timeHeld * timeToRace
		if dist > recordDist {
			waysToWin++
		}
		if dist < lastDist && dist < recordDist {
			// no longer increasing, stop processing
			break
		}
		lastDist = dist
		i++
	}
	log.Printf("ways: %d, iterations: %d", waysToWin, i)
	return nil
}
