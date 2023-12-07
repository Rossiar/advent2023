package main

import (
	"fmt"
	"slices"
	"strconv"
	"unicode"

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

type Num struct {
	Gear  string
	Value string
	X, Y  int
}

func task(lines []string) error {
	total := 0
	for y, line := range lines {
		for x, char := range line {
			if char != '*' {
				continue
			}
			fmt.Printf("found gear %c at (%d, %d)\n", char, x, y)
			nums, err := findNumbers(lines, x, y)
			if err != nil {
				return err
			}
			uniqueNums := slices.Compact(nums)
			if len(uniqueNums) != 2 {
				continue
			}
			ratio := 1
			for _, num := range uniqueNums {
				val, err := strconv.Atoi(num.Value)
				if err != nil {
					return fmt.Errorf("(%d,%d): %w", num.X, num.Y, err)
				}
				ratio *= val
			}
			total += ratio
		}
	}

	fmt.Printf("total: %d\n", total)
	return nil
}

func findNumbers(schematic []string, x, y int) ([]Num, error) {
	words := make([]Num, 0)
	// north west
	if y != 0 && x != 0 && unicode.IsDigit(rune(schematic[y-1][x-1])) {
		word := findNumber(x-1, schematic[y-1])
		word.Y = y - 1
		words = append(words, word)
	}
	// north
	if y != 0 && unicode.IsDigit(rune(schematic[y-1][x])) {
		word := findNumber(x, schematic[y-1])
		word.Y = y - 1
		words = append(words, word)
	}
	// north east
	if y != 0 && x != len(schematic[y-1])-1 && unicode.IsDigit(rune(schematic[y-1][x+1])) {
		word := findNumber(x+1, schematic[y-1])
		word.Y = y - 1
		words = append(words, word)
	}
	// west
	if x != 0 && unicode.IsDigit(rune(schematic[y][x-1])) {
		word := findNumber(x-1, schematic[y])
		word.Y = y
		words = append(words, word)
	}
	// east
	if x != len(schematic[y])-1 && unicode.IsDigit(rune(schematic[y][x+1])) {
		word := findNumber(x+1, schematic[y])
		word.Y = y
		words = append(words, word)
	}
	// south west
	if y != len(schematic)-1 && x != 0 && unicode.IsDigit(rune(schematic[y+1][x-1])) {
		word := findNumber(x-1, schematic[y+1])
		word.Y = y + 1
		words = append(words, word)
	}
	// south
	if y != len(schematic)-1 && unicode.IsDigit(rune(schematic[y+1][x])) {
		word := findNumber(x, schematic[y+1])
		word.Y = y + 1
		words = append(words, word)
	}
	// south east
	if y != len(schematic)-1 && x != len(schematic[y+1])-1 && unicode.IsDigit(rune(schematic[y+1][x+1])) {
		word := findNumber(x+1, schematic[y+1])
		word.Y = y + 1
		words = append(words, word)
	}
	return words, nil
}

func findNumber(x int, line string) Num {
	var start int
	word := line[x : x+1]
	// read left
	for i := x - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(line[i])) {
			start = i + 1
			break
		}
		word = string(line[i]) + word
	}
	// read right
	for i := x + 1; i < len(line); i++ {
		if !unicode.IsDigit(rune(line[i])) {
			break

		}
		word += string(line[i])
	}
	return Num{
		Value: word,
		X:     start,
	}
}
