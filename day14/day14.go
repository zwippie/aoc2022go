package day14

import (
	"bufio"
	"bytes"
	"fmt"
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

var minX, minY = math.MaxInt, math.MaxInt
var maxX, maxY = math.MinInt, math.MinInt

func PartA(input []byte) any {
	cave := parseInput(input)
	fmt.Printf("cave: %v\n", cave)
	cave.addSand(Pos{500, 0})
	fmt.Printf("cave: %v\n", cave)

	return 0
}

func PartB(input []byte) any {
	return 0
}

func (cave Cave) addSand(fromPos Pos) bool {
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
			minX = min(minX, x)
			maxX = max(maxX, x)
			minY = min(minY, y)
			maxY = max(maxY, y)
			segments = append(segments, Pos{x, y})
		}
		for i := 0; i < len(segments)-1; i++ {
			cave = placeRocks(cave, segments[i], segments[i+1])
		}
	}
	fmt.Println(minX, maxX, minY, maxY)
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
