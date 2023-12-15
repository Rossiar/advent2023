package main

import (
	"errors"
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

type direction string

const (
	_     = iota
	North = "north"
	East  = "east"
	South = "south"
	West  = "west"
)

type point struct {
	X, Y int
}

func (p point) Transform(dir direction) point {
	switch dir {
	case North:
		return point{p.X, p.Y - 1}
	case East:
		return point{p.X + 1, p.Y}
	case South:
		return point{p.X, p.Y + 1}
	case West:
		return point{p.X - 1, p.Y}
	}
	panic(fmt.Errorf("unknown direction %s", dir))
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
func Task1(lines []string) error {
	var startX, startY int
	for y, line := range lines {
		if x := strings.Index(line, "S"); x != -1 {
			startX = x
			startY = y
			break
		}
	}
	log.Println(start(lines, startX, startY))
	return nil
}

func start(lines []string, x, y int) int {
	start := point{x, y}
	log.Println("looking north from S")
	nSteps, err := follow(lines, North, start.Transform(North), 1)
	if err == nil {
		return nSteps / 2
	}
	log.Println("looking east from S")
	eSteps, err := follow(lines, East, start.Transform(East), 1)
	if err == nil {
		return eSteps / 2
	}
	log.Println("looking south from S")
	sSteps, err := follow(lines, South, start.Transform(South), 1)
	if err == nil {
		return sSteps / 2
	}
	log.Println("looking west from S")
	wSteps, err := follow(lines, West, start.Transform(West), 1)
	if err == nil {
		return wSteps / 2
	}
	return -1
}

var ErrDeadEnd = errors.New("dead end")

func follow(grid []string, dir direction, p point, steps int) (int, error) {
	if p.Y < 0 && dir == North {
		log.Printf("dead end going %s at %+v", dir, p)
		return -1, ErrDeadEnd
	}
	if p.X < 0 && dir == West {
		log.Printf("dead end going %s at %+v", dir, p)
		return -1, ErrDeadEnd
	}
	if p.Y > len(grid)-1 && dir == South {
		log.Printf("dead end going %s at %+v", dir, p)
		return -1, ErrDeadEnd
	}
	if p.X > len(grid[p.Y])-1 && dir == East {
		log.Printf("dead end going %s at %+v", dir, p)
		return -1, ErrDeadEnd
	}
	//log.Printf("going %s to %+v", dir, p)
	switch grid[p.Y][p.X] {
	case 'S':
		return steps, nil
	case '|':
		if dir == North || dir == South {
			return follow(grid, dir, p.Transform(dir), steps+1)
		}
	case '-':
		if dir == East || dir == West {
			return follow(grid, dir, p.Transform(dir), steps+1)
		}
	case 'L':
		if dir == South {
			return follow(grid, East, p.Transform(East), steps+1)
		} else if dir == West {
			return follow(grid, North, p.Transform(North), steps+1)
		}
	case 'J':
		if dir == East {
			return follow(grid, North, p.Transform(North), steps+1)
		} else if dir == South {
			return follow(grid, West, p.Transform(West), steps+1)
		}
	case '7':
		if dir == North {
			return follow(grid, West, p.Transform(West), steps+1)
		} else if dir == East {
			return follow(grid, South, p.Transform(South), steps+1)
		}
	case 'F':
		if dir == North {
			return follow(grid, East, p.Transform(East), steps+1)
		} else if dir == West {
			return follow(grid, South, p.Transform(South), steps+1)
		}
	}
	log.Printf("dead end at %c going %s at %+v", grid[p.Y][p.X], dir, p)
	return -1, ErrDeadEnd
}

func Task2(lines []string) error {
	return nil
}
