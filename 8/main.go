package main

import (
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
	if err := Task1(lines); err != nil {
		panic(err.Error())
	}
}

type node struct {
	ID, Left, Right string
}

func Task1(lines []string) error {
	directions := lines[0]
	nodes := make(map[string]node, 0)
	for i := 2; i < len(lines); i++ {
		lineParts := strings.Split(lines[i], "=")
		id := strings.TrimSpace(lineParts[0])
		nodeParts := strings.Split(lineParts[1], ",")
		nodes[id] = node{
			ID:    id,
			Left:  nodeParts[0][2:],
			Right: nodeParts[1][1:4],
		}
	}
	current := nodes["AAA"]
	dest := "ZZZ"
	steps := 0
	d := 0
	for {
		direction := directions[d]
		steps++
		if direction == 'L' {
			if current.Left == dest {
				break
			}
			current = nodes[current.Left]
		} else {
			if current.Right == dest {
				break
			}
			current = nodes[current.Right]
		}
		d++
		if d >= len(directions) {
			d = 0
		}
	}
	log.Println(steps)
	return nil
}
