package main

import (
	"fmt"
	"log"
	"os"
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
	expanded, galaxies := expand(lines)
	for _, line := range expanded {
		log.Println(line)
	}
	log.Println(galaxies)
	total := 0
	pairs := make([]pair, 0)
	for self, galaxy := range galaxies {
		for i := self + 1; i < len(galaxies); i++ {
			target := galaxies[i]
			dist := diff(galaxy.X, target.X) + diff(galaxy.Y, target.Y)
			total += dist
			pairs = append(pairs, pair{fmt.Sprintf("%d->%d", self+1, i+1), galaxy, target, dist})
		}
	}
	log.Println(pairs)
	//log.Println(pairs)
	log.Println(total)
	return nil
}

func diff(a, b int) int {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}

type pair struct {
	Name     string
	A, B     point
	Distance int
}

type point struct {
	X, Y int
}

func expand(lines []string) ([]string, []point) {
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "#") {
			continue
		}
		// expand row
		lines = append(lines[:i+1], lines[i:]...)
		lines[i] = line
		i++
	}
	for i := 0; i < len(lines[0]); i++ {
		hasGalaxy := false
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '#' {
				hasGalaxy = true
				break
			}
		}
		if hasGalaxy {
			continue
		}
		// expand column
		for j := 0; j < len(lines); j++ {
			line := []rune(lines[j])
			line = append(line[:i+1], line[i:]...)
			line[i] = '.'
			lines[j] = string(line)
		}
		i++
	}
	galaxies := make([]point, 0)
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				galaxies = append(galaxies, point{x, y})
			}
		}
	}
	return lines, galaxies
}

func Task2(lines []string) error {
	return nil
}
