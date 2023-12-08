package advent2023

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLinesFromFile(name string) ([]string, error) {
	raw, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	scan := bufio.NewScanner(raw)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines, nil
}

func ReadIntsFromString(src string) ([]int, error) {
	ss := strings.Fields(src)
	nums := make([]int, len(ss))
	for i, s := range ss {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return nums, nil
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
