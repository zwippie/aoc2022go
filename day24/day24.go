package day24

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"strings"
)

type Direction uint8

const (
	North Direction = iota
	East
	South
	West
)

type Pos struct {
	row, col int
}

type Node struct {
	pos       Pos
	blizzards Blizzards
	wall      bool
}

type Cave map[Pos]*Node

type Blizzard struct {
	pos          Pos
	direction    Direction
	wrapToRowCol int
}

type Blizzards []*Blizzard

var maxRow, maxCol int
var entry, exit Pos

func PartA(input []byte) any {
	cave, blizzards := parseInput(input)
	fmt.Println(cave)

	for minute := 0; minute < 18; minute++ {
		blizzards = moveBlizzards(cave, blizzards)
		cave = updateCave(cave, blizzards)
		fmt.Println(cave)
	}

	return 0
}

func PartB(input []byte) any {
	return 0
}

func moveBlizzards(cave Cave, blizzards Blizzards) Blizzards {
	for _, blizzard := range blizzards {
		np := nextPos(blizzard.pos, blizzard.direction)
		if cave[np].wall { // wrap around
			switch blizzard.direction {
			case North:
				blizzard.pos = Pos{blizzard.wrapToRowCol, blizzard.pos.col}
			case East:
				blizzard.pos = Pos{blizzard.pos.row, blizzard.wrapToRowCol}
			case South:
				blizzard.pos = Pos{blizzard.wrapToRowCol, blizzard.pos.col}
			case West:
				blizzard.pos = Pos{blizzard.pos.row, blizzard.wrapToRowCol}
			}
		} else {
			switch blizzard.direction {
			case North:
				blizzard.pos = Pos{blizzard.pos.row - 1, blizzard.pos.col}
			case East:
				blizzard.pos = Pos{blizzard.pos.row, blizzard.pos.col + 1}
			case South:
				blizzard.pos = Pos{blizzard.pos.row + 1, blizzard.pos.col}
			case West:
				blizzard.pos = Pos{blizzard.pos.row, blizzard.pos.col - 1}
			}
		}
	}
	return blizzards
}

func updateCave(cave Cave, blizzards Blizzards) Cave {
	for _, node := range cave {
		node.blizzards = []*Blizzard{}
	}
	for _, blizzard := range blizzards {
		cave[blizzard.pos].blizzards = append(cave[blizzard.pos].blizzards, blizzard)
	}
	return cave
}

func nextPos(pos Pos, direction Direction) Pos {
	switch direction {
	case North:
		return Pos{pos.row - 1, pos.col}
	case East:
		return Pos{pos.row, pos.col + 1}
	case South:
		return Pos{pos.row + 1, pos.col}
	case West:
		return Pos{pos.row, pos.col - 1}
	}
	return Pos{} // unreachable
}

func (cave Cave) String() string {
	result := ""
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			result += cave[Pos{row, col}].String()
		}
		result += "\n"
	}
	return result
}

func (node *Node) String() string {
	if node.pos == entry {
		return "e"
	} else if node.pos == exit {
		return "x"
	} else if node.wall {
		return "#"
	} else if len(node.blizzards) > 1 {
		return fmt.Sprint(len(node.blizzards))
	} else if len(node.blizzards) == 1 {
		return node.blizzards[0].direction.String()
	} else {
		return "."
	}
}

func (d Direction) String() string {
	switch d {
	case North:
		return "^"
	case East:
		return ">"
	case South:
		return "v"
	case West:
		return "<"
	}
	return "?"
}

func parseInput(input []byte) (Cave, Blizzards) {
	cave := make(Cave)
	blizzards := Blizzards{}
	maxRow, maxCol = math.MinInt, math.MinInt
	for row, line := range strings.Split(string(input), "\n") {
		maxRow = utils.Max(maxRow, row)
		for col, value := range strings.Split(line, "") {
			maxCol = utils.Max(maxCol, col)
			node := &Node{pos: Pos{row, col}}
			if value == "#" {
				node.wall = true
			} else if value == "^" {
				blizzard := &Blizzard{pos: Pos{row, col}, direction: North}
				blizzards = append(blizzards, blizzard)
				node.blizzards = append(node.blizzards, blizzard)
			} else if value == ">" {
				blizzard := &Blizzard{pos: Pos{row, col}, direction: East}
				blizzards = append(blizzards, blizzard)
				node.blizzards = append(node.blizzards, blizzard)
			} else if value == "v" {
				blizzard := &Blizzard{pos: Pos{row, col}, direction: South}
				blizzards = append(blizzards, blizzard)
				node.blizzards = append(node.blizzards, blizzard)
			} else if value == "<" {
				blizzard := &Blizzard{pos: Pos{row, col}, direction: West}
				blizzards = append(blizzards, blizzard)
				node.blizzards = append(node.blizzards, blizzard)
			}
			cave[node.pos] = node
		}
	}

	for _, blizzard := range blizzards {
		switch blizzard.direction {
		case North:
			blizzard.wrapToRowCol = maxRow - 1
		case East:
			blizzard.wrapToRowCol = 1
		case South:
			blizzard.wrapToRowCol = 1
		case West:
			blizzard.wrapToRowCol = maxCol - 1
		}
	}

	entry = Pos{0, 1}
	exit = Pos{maxRow, maxCol - 1}

	return cave, blizzards
}
