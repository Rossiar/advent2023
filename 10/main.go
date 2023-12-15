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
	if err := Task2(lines); err != nil {
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
	path, err := start(lines)
	if err != nil {
		return err
	}
	log.Println(len(path) / 2)
	return nil
}

func start(lines []string) ([]point, error) {
	start := point{}
	for y, line := range lines {
		if x := strings.Index(line, "S"); x != -1 {
			start.X = x
			start.Y = y
			break
		}
	}
	path := []point{start}
	log.Println("looking north from S")
	nSteps, err := traverse(lines, North, start.Transform(North), path)
	if err == nil {
		return nSteps, nil
	}
	log.Println("looking east from S")
	eSteps, err := traverse(lines, East, start.Transform(East), path)
	if err == nil {
		return eSteps, nil
	}
	log.Println("looking south from S")
	sSteps, err := traverse(lines, South, start.Transform(South), path)
	if err == nil {
		return sSteps, nil
	}
	log.Println("looking west from S")
	wSteps, err := traverse(lines, West, start.Transform(West), path)
	if err == nil {
		return wSteps, nil
	}
	return nil, fmt.Errorf("no paths from %+v", start)
}

var ErrDeadEnd = errors.New("dead end")

func traverse(grid []string, dir direction, p point, path []point) ([]point, error) {
	if p.Y < 0 && dir == North {
		log.Printf("dead end going %s at %+v", dir, p)
		return nil, ErrDeadEnd
	}
	if p.X < 0 && dir == West {
		log.Printf("dead end going %s at %+v", dir, p)
		return nil, ErrDeadEnd
	}
	if p.Y > len(grid)-1 && dir == South {
		log.Printf("dead end going %s at %+v", dir, p)
		return nil, ErrDeadEnd
	}
	if p.X > len(grid[p.Y])-1 && dir == East {
		log.Printf("dead end going %s at %+v", dir, p)
		return nil, ErrDeadEnd
	}
	//log.Printf("going %s to %+v", dir, p)
	switch grid[p.Y][p.X] {
	case 'S':
		return path, nil
	case '|':
		if dir == North || dir == South {
			return traverse(grid, dir, p.Transform(dir), append(path, p))
		}
	case '-':
		if dir == East || dir == West {
			return traverse(grid, dir, p.Transform(dir), append(path, p))
		}
	case 'L':
		if dir == South {
			return traverse(grid, East, p.Transform(East), append(path, p))
		} else if dir == West {
			return traverse(grid, North, p.Transform(North), append(path, p))
		}
	case 'J':
		if dir == East {
			return traverse(grid, North, p.Transform(North), append(path, p))
		} else if dir == South {
			return traverse(grid, West, p.Transform(West), append(path, p))
		}
	case '7':
		if dir == North {
			return traverse(grid, West, p.Transform(West), append(path, p))
		} else if dir == East {
			return traverse(grid, South, p.Transform(South), append(path, p))
		}
	case 'F':
		if dir == North {
			return traverse(grid, East, p.Transform(East), append(path, p))
		} else if dir == West {
			return traverse(grid, South, p.Transform(South), append(path, p))
		}
	}
	log.Printf("dead end at %c going %s at %+v", grid[p.Y][p.X], dir, p)
	return nil, ErrDeadEnd
}

func Task2(lines []string) error {
	path, err := start(lines)
	if err != nil {
		return err
	}
	log.Println(len(path))

	// shoelace algorithm
	trailing := 0
	for i := 0; i < len(path); i++ {
		if i == len(path)-1 {
			trailing += (path[i].Y + path[0].Y) * (path[i].X - path[0].X)
		} else {
			trailing += (path[i].Y + path[i+1].Y) * (path[i].X - path[i+1].X)
		}
	}
	area := trailing / 2

	// Pick's theorem
	inner := area - (len(path) / 2) + 1
	log.Println(inner)
	return nil
}
