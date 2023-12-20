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
	if err := Task2(lines); err != nil {
		panic(err.Error())
	}
}

func Task1(lines []string) error {
	galaxies := expand(lines, 1)
	log.Println(len(galaxies))
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
	log.Println(len(pairs))
	//log.Println(pairs)
	log.Println(total)
	return nil
}

func Task2(lines []string) error {
	galaxies := expand(lines, 1000000)
	log.Println(len(galaxies))
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
	log.Println(len(pairs))
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

func expand(lines []string, expandBy int) []point {
	if expandBy != 1 {
		expandBy--
	}
	modX := 0
	modY := 0
	galaxies := make([]point, 0)
	for y := 0; y < len(lines); y++ {
		if !strings.Contains(lines[y], "#") {
			modY += expandBy
			continue
		}
		for x := 0; x < len(lines[y]); x++ {
			hasGalaxy := false
			for s := 0; s < len(lines); s++ {
				if lines[s][x] == '#' {
					hasGalaxy = true
					break
				}
			}
			if !hasGalaxy {
				modX += expandBy
				continue
			}
			if lines[y][x] == '#' {
				galaxies = append(galaxies, point{
					X: x + modX,
					Y: y + modY,
				})
			}
		}
		modX = 0
	}
	return galaxies
}
