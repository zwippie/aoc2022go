package day17

import (
	"aoc2022/utils"
	"fmt"
)

// type Shape [][]bool

type Jets struct {
	jets    []int
	current int
}

type Shape struct {
	bottoms []int
	tops    [][]int // if landing on shapeCol, how much to add to landing row
}
type Shapes struct {
	shapes  []Shape
	current int
}

func PartA(input []byte) any {
	jets := parseInput(input)
	shapes := createShapes()
	fmt.Printf("jets: %v\n", jets)
	fmt.Printf("shapes: %v\n", shapes)

	floors := make([]int, 7)

	for i := 0; i < 2022; i++ { // 2022
		fmt.Printf("i: %v\n", i)
		maxFloor := utils.MaxIn(floors)
		shape := shapes.next()
		col := 2
		row := maxFloor + 3

		for { // shift and drop
			shift := jets.next()
			col = moveShape(floors, shape, col, row, shift)
			if shapeCol, landed := hasLanded(floors, shape, col, row); landed {
				// adjust floors
				for c, top := range shape.tops[shapeCol] {
					floors[col+c] = row + top
				}
				break
			} else {
				row -= 1
			}
		}
	}

	return utils.MaxIn(floors)
}

func PartB(input []byte) any {
	return 0
}

// return column of landing, -1 if not landed
func hasLanded(floors []int, shape Shape, col int, row int) (shapeCol int, landed bool) {
	for shapeCol, hitRow := range shape.bottoms {
		if row+hitRow == floors[col+shapeCol] {
			return shapeCol, true
		}
	}
	return 0, false
}

// return new col the shape is on after moving
func moveShape(floors []int, shape Shape, col int, row int, shift int) int {
	if shift == 1 {
		if col+len(shape.bottoms) < len(floors) { // no hit right border?
			return col + 1
		}
	} else if shift == -1 {
		if col > 0 {
			return col - 1
		}
	}

	return col // no movement
}

func (jets *Jets) next() int {
	jet := jets.jets[jets.current]
	jets.current = (jets.current + 1) % len(jets.jets)
	return jet
}

func createShapes() *Shapes {
	return &Shapes{
		shapes: []Shape{
			{bottoms: []int{0, 0, 0, 0}, tops: [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}}, //, []int{1, 1, 1, 1}},
			{bottoms: []int{1, 0, 1}, tops: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 2, 1}}},
			{bottoms: []int{0, 0, 0}, tops: [][]int{{1, 1, 3}, {1, 1, 3}, {1, 1, 3}}},
			{bottoms: []int{0}, tops: [][]int{{4}}},
			{bottoms: []int{0, 0}, tops: [][]int{{2, 2}, {2, 2}}},
		},
		current: 0,
	}
}

func (shapes *Shapes) next() Shape {
	shape := shapes.shapes[shapes.current]
	shapes.current = (shapes.current + 1) % len(shapes.shapes)
	return shape
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
