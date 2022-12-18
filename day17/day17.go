package day17

import "fmt"

type Jets struct {
	jets    []int
	current int
}

// 3109
func PartA(input []byte) any {
	jets := parseInput(input)
	shapes := GetShapes()
	grid := NewGrid()

	grid.buildTower(shapes, jets, 2022)

	return grid.maxRow
}

func PartB(input []byte) any {
	jets := parseInput(input)
	shapes := GetShapes()
	grid := NewGrid()

	grid.buildTower(shapes, jets, 10_000)
	// grid.Print()

	return grid.maxRow
}

func (grid *Grid) buildTower(shapes *Shapes, jets *Jets, steps int) {
	shapeCount := 0
	for i := 0; i < steps; i++ {
		shape := shapes.Next()
		shapeCount++
		pos := Pos{grid.maxRow + 4, 2}

		for { // shift and drop
			shift := jets.next()
			if grid.HasRoom(shape, Pos{pos.row, pos.col + shift}) {
				pos = Pos{pos.row, pos.col + shift}
			}
			if grid.HasRoom(shape, Pos{pos.row - 1, pos.col}) {
				pos = Pos{pos.row - 1, pos.col}
			} else {
				grid.PlaceShape(shape, pos)
				break
			}
		}
		if (i % 1_000_000) == 0 {
			fmt.Printf("i: %v, minRow: %v, maxRow: %v, diff: %v\n", i, grid.minRow, grid.maxRow, grid.maxRow-grid.minRow)
		}
	}
}

func (jets *Jets) next() int {
	jet := jets.jets[jets.current]
	jets.current = (jets.current + 1) % len(jets.jets)
	return jet
}

func parseInput(input []byte) *Jets {
	result := []int{}
	for _, b := range input {
		switch b {
		case 60: // <
			result = append(result, -1)
		case 62: // >
			result = append(result, 1)
		}
	}
	return &Jets{result, 0}
}
