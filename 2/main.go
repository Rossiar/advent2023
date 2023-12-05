package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	if err := do(); err != nil {
		panic(err.Error())
	}
}

func do() error {
	lines, err := aoc.ReadLinesFromFile("in.txt")
	if err != nil {
		return err
	}
	total := 0
	for i, line := range lines {
		var maxBlue, maxGreen, maxRed int
		var powers int
		line = strings.TrimPrefix(line, "Game ")
		lineParts := strings.Split(line, ":")
		game := lineParts[1]
		gameNum, err := strconv.Atoi(lineParts[0])
		if err != nil {
			return fmt.Errorf("invalid game number on line %d: %w", i+1, err)
		}
		sets := strings.Split(game, ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				cubeParts := strings.Split(cube, " ")
				num, err := strconv.Atoi(cubeParts[0])
				if err != nil {
					return fmt.Errorf("game on line %d had invalid set %s: %w", i+1, cube, err)
				}
				switch cubeParts[1] {
				case "green":
					if maxGreen < num {
						maxGreen = num
					}
				case "blue":
					if maxBlue < num {
						maxBlue = num
					}
				case "red":
					if maxRed < num {
						maxRed = num
					}
				}
			}
		}
		powers = maxRed * maxBlue * maxGreen
		fmt.Printf("Game %d: %d * %d * %d = %d\n", gameNum, maxRed, maxGreen, maxBlue, powers)
		total += powers
	}
	fmt.Println(total)

	return nil
}
