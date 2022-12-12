package day12

import (
	"fmt"
	"math"
	"strings"
)

type Grid [][]byte
type Pos struct {
	row int
	col int
}
type Path []Pos

var rows int
var cols int

func PartA(input []byte) any {
	grid, start := parseInput(input)
	// fmt.Printf("grid: %v\n", grid)
	fmt.Printf("start: %v\n", start)
	routes := []Path{}
	findPath(grid, start, []Pos{}, &routes)
	fmt.Println(len(routes), "routes found")

	minLen := math.MaxInt
	for _, route := range routes {
		if len(route) < minLen {
			minLen = len(route)
		}
	}
	return minLen
}

func findPath(grid Grid, pos Pos, trail Path, routes *[]Path) {
	if len(trail) > 300 || discardTrail(trail, *routes) {
		fmt.Println("discarding trail")
		return
	}
	elevation := grid[pos.row][pos.col]
	// fmt.Printf("pos %v, elevation: %v\n", pos, elevation)
	for _, apos := range adjacent(pos) {
		if inTrail(trail, apos) {
			// fmt.Println(apos, "already visited in", trail)
			continue
		}
		if elevation == 122 && grid[apos.row][apos.col] == 69 { // E
			trail = append(trail, apos)
			fmt.Println("found a route!", len(trail))
			*routes = append(*routes, trail)
		}
		if grid[apos.row][apos.col] == elevation || grid[apos.row][apos.col] == elevation+1 {
			// fmt.Println("at", pos, elevation, "going to", apos, grid[apos.row][apos.col], "trail", trail)
			findPath(grid, apos, append(trail, pos), routes)
		}
	}
}

func adjacent(pos Pos) []Pos {
	results := []Pos{}
	if pos.row > 0 {
		results = append(results, Pos{pos.row - 1, pos.col})
	}
	if pos.row < rows-1 {
		results = append(results, Pos{pos.row + 1, pos.col})
	}
	if pos.col > 0 {
		results = append(results, Pos{pos.row, pos.col - 1})
	}
	if pos.col < cols-1 {
		results = append(results, Pos{pos.row, pos.col + 1})
	}
	return results
}

func inTrail(trail []Pos, pos Pos) bool {
	for _, p := range trail {
		if p == pos {
			return true
		}
	}
	return false
}

func discardTrail(trail []Pos, routes []Path) bool {
	for _, route := range routes {
		if len(trail) > len(route) {
			return true
		}
	}
	return false
}

func PartB(input []byte) any {
	return 0
}

func parseInput(input []byte) (Grid, Pos) {
	result := Grid{}
	start := Pos{}
	parts := strings.Split(string(input), "\n")
	for r, part := range parts {
		row := []byte{}
		for c, b := range []byte(part) {
			if b == 83 { // S
				start = Pos{r, c}
				b = 96 // one befor a
			}
			row = append(row, b)
		}
		result = append(result, row)
	}
	rows = len(result)
	cols = len(result[0])
	return result, start
}
