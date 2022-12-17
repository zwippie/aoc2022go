package day17

import (
	"aoc2022/utils"
	"fmt"
)

type Pos struct {
	row, col int
}
type Grid struct {
	minRow int
	maxRow int
	rocks  map[Pos]bool
}

type Shape []Pos

type Shapes struct {
	shapes  []Shape
	current int
}

const maxCol int = 6

func (shapes *Shapes) Next() Shape {
	shape := shapes.shapes[shapes.current]
	shapes.current = (shapes.current + 1) % len(shapes.shapes)
	return shape
}

func NewGrid() *Grid {
	grid := Grid{rocks: make(map[Pos]bool)}
	for col := 0; col <= maxCol; col++ {
		grid.rocks[Pos{0, col}] = true
	}
	return &grid
}

func (g *Grid) PlaceShape(s Shape, p Pos) {
	rowsToCheck := make(map[int]bool)
	maxRow := 0
	// place shape
	for _, pos := range s {
		pt := p.Add(pos)
		g.rocks[pt] = true
		rowsToCheck[pt.row] = true
		maxRow = utils.Max(maxRow, pt.row)
	}
	g.maxRow = utils.Max(g.maxRow, maxRow)
	// full line created?
	// for row := range rowsToCheck {
	// 	fullRow := true
	// 	for col := 0; col <= maxCol; col++ {
	// 		if !g.rocks[Pos{row, col}] {
	// 			fullRow = false
	// 			break
	// 		}
	// 	}
	// 	if fullRow {
	// 		// remove everything below the fullRow to free memory
	// 		// fmt.Println("full row at", row, "prev was", g.minRow)
	// 		for row2 := g.minRow; row2 < row; row2++ {
	// 			for col2 := 0; col2 <= maxCol; col2++ {
	// 				delete(g.rocks, Pos{row2, col2})
	// 			}
	// 		}
	// 		g.minRow = row
	// 		break
	// 	}
	// }
}

func (g *Grid) HasRoom(s Shape, p Pos) bool {
	for _, pos := range s {
		pt := p.Add(pos)
		if pt.col < 0 || pt.col > maxCol || g.rocks[pt] {
			return false
		}
	}
	return true
}

func (g *Grid) Print() {
	for row := g.minRow; row <= g.maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if g.rocks[Pos{row, col}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (p1 Pos) Add(p2 Pos) Pos {
	return Pos{p1.row + p2.row, p1.col + p2.col}
}

func GetShapes() *Shapes {
	shapes := []Shape{
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	}
	return &Shapes{shapes, 0}
}
