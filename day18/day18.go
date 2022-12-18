package day18

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

type Cube struct {
	x, y, z int
}

type Grid map[Cube]int

var minX, minY, minZ, maxX, maxY, maxZ = 100, 100, 100, 0, 0, 0

// 4636
func PartA(input []byte) any {
	data := parseInput(input)
	grid := make(Grid)
	// fmt.Printf("data: %v\n", data)

	for _, c := range data {
		grid = AddToGrid(grid, c)
	}

	return SidesTotal(grid)
}

// 2564 too low
func PartB(input []byte) any {
	data := parseInput(input)
	grid := make(Grid)
	// fmt.Printf("data: %v\n", data)

	for _, c := range data {
		minX, minY, minZ = utils.Min(minX, c.x), utils.Min(minY, c.y), utils.Min(minZ, c.z)
		maxX, maxY, maxZ = utils.Max(maxX, c.x), utils.Max(maxY, c.y), utils.Max(maxZ, c.z)
		grid = AddToGrid(grid, c)
	}
	fmt.Println("grid size", len(grid))
	fmt.Println("min max", minX, maxX, minY, maxY, minZ, maxZ)
	volume := (maxX - minX + 1) * (maxY - minY + 1) * (maxZ - minZ + 1)
	fmt.Printf("volume: %v\n", volume)

	// candidate pockets
	candidates := make(map[Cube]bool)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				if _, ok := grid[Cube{x, y, z}]; !ok {
					candidates[Cube{x, y, z}] = true
				}
			}
		}
	}
	fmt.Printf("len(candidates): %v\n", len(candidates))

	// candidate pockets that are on the border of the cube cannot be pockets
	borders := make(map[Cube]bool)
	for c := range candidates {
		if OnGridBorder(grid, c) {
			borders[c] = true
			delete(candidates, c)
		}
	}
	fmt.Printf("len(borders): %v\n", len(borders))
	fmt.Printf("len(candidates): %v\n", len(candidates))

	// check all borders to see if they are next to a candidate. Remove candidate if so.
	changes := 1
	for changes > 0 {
		borders, candidates, changes = AddBorders(grid, borders, candidates)
	}
	fmt.Printf("after reduce len(borders): %v\n", len(borders))
	fmt.Printf("after reduce len(candidates): %v\n", len(candidates))

	// candidates now all should be part of a pocket
	// if a grid cell is next to a pocket, it has one exposed surface less
	for c := range candidates {
		for _, s := range surrounding(c) {
			if _, ok := grid[s]; ok {
				grid[s]--
			}
		}
	}

	return SidesTotal(grid)
}

func AddBorders(grid Grid, borders map[Cube]bool, candidates map[Cube]bool) (map[Cube]bool, map[Cube]bool, int) {
	changes := 0
	for b := range borders {
		for _, c := range surrounding(b) {
			if _, ok := candidates[c]; ok {
				// found a neighbour in candidates next to a border
				delete(candidates, c)
				borders[c] = true
				changes++
			}
		}
	}
	return borders, candidates, changes
}

func OnGridBorder(grid Grid, cube Cube) bool {
	return cube.x == minX || cube.x == maxX || cube.y == minY || cube.y == maxY || cube.z < minZ || cube.z == maxZ
}

func OutsideGrid(grid Grid, cube Cube) bool {
	return cube.x < minX || cube.x > maxX || cube.y < minY || cube.y > maxY || cube.z < minZ || cube.z > maxZ
}

func SidesTotal(grid Grid) int {
	result := 0
	for _, sides := range grid {
		result += sides
	}
	return result
}

func OutsideSidesTotal(grid Grid) int {
	return 0
}

func NewCube(x, y, z int) Cube {
	return Cube{x, y, z}
}

func AddToGrid(grid Grid, cube Cube) Grid {
	grid[cube] = 6
	for _, c := range surrounding(cube) {
		if _, ok := grid[c]; ok {
			grid[c]--
			grid[cube]--
		}
	}
	return grid
}

func surrounding(c Cube) []Cube {
	return []Cube{
		{c.x - 1, c.y, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y, c.z - 1},
		{c.x + 1, c.y, c.z},
		{c.x, c.y + 1, c.z},
		{c.x, c.y, c.z + 1},
	}
}

func parseInput(input []byte) []Cube {
	result := []Cube{}
	for _, line := range strings.Split(string(input), "\n") {
		// fmt.Printf("line: %v\n", line)
		parts := strings.Split(line, ",")
		cube := Cube{utils.ParseInt(parts[0]), utils.ParseInt(parts[1]), utils.ParseInt(parts[2])}
		result = append(result, cube)
	}
	return result
}
