package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	aoc "github.com/rossiar/advent2023"
)

func main() {
	if err := do(); err != nil {
		panic(err.Error())
	}
}

func do() error {
	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	revReplacer := strings.NewReplacer(
		"eno", "1",
		"owt", "2",
		"eerht", "3",
		"ruof", "4",
		"evif", "5",
		"xis", "6",
		"neves", "7",
		"thgie", "8",
		"enin", "9",
	)
	lines, err := aoc.ReadLinesFromFile("in.txt")
	if err != nil {
		return err
	}
	total := 0
	for i, line := range lines {
		sanitized := replacer.Replace(line)
		first := findFirstDigit(sanitized)
		rev := aoc.Reverse(line)
		revSanitized := revReplacer.Replace(rev)
		second := findFirstDigit(revSanitized)
		rawNum := string(first) + string(second)
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			return fmt.Errorf("%s on line %d was not a number: %w", rawNum, i, err)
		}
		fmt.Println(num)
		total += num
	}
	fmt.Println(total)
	return nil
}

func findFirstDigit(str string) rune {
	for _, char := range str {
		if unicode.IsDigit(char) {
			return char
		}
	}
	return 0
}
