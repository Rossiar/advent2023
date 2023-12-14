package main

import (
	"fmt"
	"log"
	"os"
	"slices"

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
	total := 0
	for li, line := range lines {
		nums, err := aoc.ReadIntsFromString(line)
		if err != nil {
			return fmt.Errorf("error on line %d: %w", li, err)
		}
		inputs := [][]int{nums}
		for {
			input := diffs(inputs[len(inputs)-1])
			inputs = append(inputs, input)
			empty := make([]int, len(input))
			if slices.Equal(input, empty) {
				break
			}
		}
		for i := len(inputs) - 2; i >= 0; i-- {
			current := inputs[i]
			prev := inputs[i+1]
			if len(prev) == 0 {
				continue
			}
			value := current[len(current)-1] + prev[len(prev)-1]
			inputs[i] = append(inputs[i], value)
			if i == 0 {
				total += value
			}
		}
	}
	log.Println(total)
	return nil
}

func diffs(nums []int) []int {
	diffs := make([]int, 0)
	for k := range nums {
		if k != 0 {
			diffs = append(diffs, nums[k]-nums[k-1])
		}
	}
	return diffs
}

func Task2(lines []string) error {
	return nil
}
