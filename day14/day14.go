package day14

import (
	"bufio"
	"bytes"
	"math"
	"strconv"
	"strings"
)

type Content uint8

const (
	Empty Content = iota
	Rock
	Sand
)

type Pos struct {
	x int
	y int
}
type Cave map[Pos]Content

var maxY = math.MinInt

// 888
func PartA(input []byte) any {
	cave := parseInput(input)
	result := 0
	for !cave.addSand(Pos{500, 0}) {
		result++
	}
	return result
}

// 26461
func PartB(input []byte) any {
	cave := parseInput(input)
	maxY += 2
	cave = placeRocks(cave, Pos{500 - maxY, maxY}, Pos{500 + maxY, maxY})
	result := 0
	for !cave.addSand(Pos{500, 0}) {
		result++
		if cave[Pos{500, 0}] == Sand {
			break
		}
	}
	return result
}

func (cave Cave) addSand(pos Pos) (done bool) {
	for {
		if pos.y > maxY {
			return true
		} else if cave[Pos{pos.x, pos.y + 1}] == Empty {
			pos = Pos{pos.x, pos.y + 1}
		} else if cave[Pos{pos.x - 1, pos.y + 1}] == Empty {
			pos = Pos{pos.x - 1, pos.y + 1}
		} else if cave[Pos{pos.x + 1, pos.y + 1}] == Empty {
			pos = Pos{pos.x + 1, pos.y + 1}
		} else {
			cave[pos] = Sand
			break
		}
	}
	return false
}

func parseInput(input []byte) Cave {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	cave := make(Cave)

	for scanner.Scan() {
		segments := []Pos{}
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		for _, part := range parts {
			coords := strings.Split(part, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			maxY = max(maxY, y)
			segments = append(segments, Pos{x, y})
		}
		for i := 0; i < len(segments)-1; i++ {
			cave = placeRocks(cave, segments[i], segments[i+1])
		}
	}
	return cave
}

func placeRocks(cave Cave, from Pos, to Pos) Cave {
	if from.x == to.x { // vertical line
		if from.y < to.y {
			for y := from.y; y <= to.y; y++ {
				cave[Pos{from.x, y}] = Rock
			}
		} else {
			for y := to.y; y <= from.y; y++ {
				cave[Pos{from.x, y}] = Rock
			}
		}
	} else if from.y == to.y { // horizontal line
		if from.x < to.x {
			for x := from.x; x <= to.x; x++ {
				cave[Pos{x, from.y}] = Rock
			}
		} else {
			for x := to.x; x <= from.x; x++ {
				cave[Pos{x, from.y}] = Rock
			}
		}
	}
	return cave

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
