package main

import (
	"log"
	"os"
	"slices"
	"strings"

	aoc "github.com/rossiar/advent2023"
	"golang.org/x/exp/maps"
)

func main() {
	filename := os.Args[1]
	lines, err := aoc.ReadLinesFromFile(filename)
	if err != nil {
		panic(err.Error())
	}
	if err := Task2(lines); err != nil {
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

func Task2(lines []string) error {
	directions := lines[0]
	nodes := make(map[string]node, 0)
	refs := make([]node, 0)
	for i := 2; i < len(lines); i++ {
		lineParts := strings.Split(lines[i], "=")
		id := strings.TrimSpace(lineParts[0])
		nodeParts := strings.Split(lineParts[1], ",")
		n := node{
			ID:    id,
			Left:  nodeParts[0][2:],
			Right: nodeParts[1][1:4],
		}
		nodes[id] = n
		if strings.HasSuffix(id, "A") {
			refs = append(refs, n)
		}
	}
	steps := 0
	d := 0
	endings := make(map[string]int, 0)
	for {
		end := false
		for i, current := range refs {
			_, found := endings[current.ID]
			if strings.HasSuffix(current.ID, "Z") && !found {
				endings[current.ID] = steps
				log.Printf("found ending %s at step %d", current.ID, steps)
				if len(refs) == 1 {
					end = true
				}
				refs = append(refs[:i], refs[i+1:]...)
			}
		}
		if end {
			log.Println(endings)
			break
		}
		direction := directions[d]
		steps++
		for i, current := range refs {
			next := current.Right
			if direction == 'L' {
				next = current.Left
			}
			//log.Printf("going %c from %s to %s", direction, current.ID, next)
			refs[i] = nodes[next]
		}
		d++
		if d == len(directions) {
			d = 0
		}
	}
	vals := maps.Values(endings)
	slices.Sort(vals)
	log.Println(LCM(vals[0], vals[1], vals[1:]...))
	return nil
}

// Greatest Common Divisor
func GCD(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

// Lowest Common Multiple
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
