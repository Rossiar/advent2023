package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	t := task{12, 13, 14}
	if err := t.do(); err != nil {
		panic(err.Error())
	}
}

type task struct {
	maxRed, maxGreen, maxBlue int
}

func (t task) do() error {
	lines, err := aoc.ReadLinesFromFile("in.txt")
	if err != nil {
		return err
	}
	total := 0
	for i, line := range lines {
		var blue, green, red int
		line = strings.TrimPrefix(line, "Game ")
		lineParts := strings.Split(line, ":")
		game := lineParts[1]
		gameNum, err := strconv.Atoi(lineParts[0])
		if err != nil {
			return fmt.Errorf("invalid game number on line %d: %w", i+1, err)
		}
		sets := strings.Split(game, ";")
		invalid := false
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				cubeParts := strings.Split(cube, " ")
				num, err := strconv.Atoi(cubeParts[0])
				if err != nil {
					return fmt.Errorf("game on line %d had invalid set %s: %w", i+1, cube, err)
				}
				if !t.validateCount(num, cubeParts[1]) {
					invalid = true
					break
				}
				switch cubeParts[1] {
				case "green":
					green += num
				case "blue":
					blue += num
				case "red":
					red += num
				}
			}
			if invalid {
				break
			}
		}
		if !invalid {
			total += gameNum
		} else {
			fmt.Println("invalid game " + strconv.Itoa(gameNum))
		}
	}
	fmt.Println(total)

	return nil
}

func (t task) validateCount(num int, colour string) bool {
	// could be a one liner but harder to read that way
	if colour == "green" && num > t.maxGreen {
		return false
	} else if colour == "blue" && num > t.maxBlue {
		return false
	} else if colour == "red" && num > t.maxRed {
		return false
	} else {
		return true
	}
}
