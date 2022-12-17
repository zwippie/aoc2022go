package day17

import "fmt"

type Pos struct {
	row, col int
}
type Grid map[Pos]bool

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
	grid := make(Grid)
	for col := 0; col <= maxCol; col++ {
		grid[Pos{0, col}] = true
	}
	return &grid
}

func (g *Grid) PlaceShape(s Shape, p Pos) {
	for _, pos := range s {
		pt := p.Add(pos)
		(*g)[pt] = true
	}
}

func (g *Grid) HasRoom(s Shape, p Pos) bool {
	for _, pos := range s {
		pt := p.Add(pos)
		if pt.col < 0 || pt.col > maxCol || (*g)[pt] {
			return false
		}
	}
	return true
}

func (g *Grid) MaxRow() int {
	maxRow := 0
	for k := range *g {
		if k.row > maxRow {
			maxRow = k.row
		}
	}
	return maxRow
}

func (g *Grid) Print() {
	maxRow := g.MaxRow()
	for row := maxRow; row > 0; row-- {
		for col := 0; col <= maxCol; col++ {
			if (*g)[Pos{row, col}] {
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
