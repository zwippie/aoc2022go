package day23

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

type PosMap[T any] map[Pos]T

type Directions []Direction

// 4075
func PartA(input []byte) any {
	elves := parseInput(input)
	for round := uint(0); round < 10; round++ {
		elves, _ = moveElves(elves, round)
	}
	return score(elves)
}

// 950 but took 21 seconds...
func PartB(input []byte) any {
	elves := parseInput(input)
	round := uint(0)
	done := false
	for {
		elves, done = moveElves(elves, round)
		if done {
			break
		}
		round++
	}
	return round + 1
}

// One round of moving elves.
func moveElves(elves PosMap[bool], round uint) (e PosMap[bool], done bool) {
	moves := make(PosMap[[]Pos]) // key = pos to move to, value = elves that want to go here
	nextElves := make(PosMap[bool])
	for pos := range elves {
		nextPos, ok := determineNextPos(elves, pos, round)
		if ok {
			moves[nextPos] = append(moves[nextPos], pos)
		} else {
			nextElves[pos] = true // elf cannot or does not want to move
		}
	}
	if len(moves) == 0 {
		return elves, true // no more move candidates, we're done
	}

	nextElves, moves = removeBlockedElves(nextElves, moves)

	// all remaining moves are to an empty spot
	for to, froms := range moves {
		if len(froms) == 1 {
			nextElves[to] = true
		} else { // too many moves to this spot, cancel them
			for _, from := range froms {
				nextElves[from] = true
			}
		}
	}
	return nextElves, false // not done yet
}

// some moves are blocked by other elves that could not move.
// keep removing them until there only valid moves left.
func removeBlockedElves(elves PosMap[bool], moves PosMap[[]Pos]) (PosMap[bool], PosMap[[]Pos]) {
	for {
		nonMovingElves := []Pos{}
		for to, froms := range moves {
			for elf := range elves {
				if elf == to { // already an elf on thi position
					nonMovingElves = append(nonMovingElves, froms...)
				}
			}
		}
		if len(nonMovingElves) == 0 {
			break
		}
		for _, elf := range nonMovingElves {
			elves[elf] = true
			delete(moves, elf)
		}
	}
	return elves, moves
}

func determineNextPos(elves PosMap[bool], pos Pos, round uint) (p Pos, ok bool) {
	isAlone := true
	for _, direction := range []Direction{North, East, South, West} {
		if !canMoveInDirection(elves, pos, direction) {
			isAlone = false
			break
		}
	}
	if !isAlone {
		for _, direction := range directionsForRound(round) {
			if canMoveInDirection(elves, pos, direction) {
				return posWhenMoving(pos, direction), true // new spot found
			}
		}
	}
	return pos, false // elf was happy here or could not find a new spot
}

func canMoveInDirection(elves PosMap[bool], pos Pos, direction Direction) bool {
	for _, p := range surroundingPositions(pos)[direction] {
		if elves[p] {
			return false
		}
	}
	return true
}

// the 3 positions when looking to N, E, S, W
func surroundingPositions(pos Pos) [4][3]Pos {
	return [4][3]Pos{
		{
			{pos.row - 1, pos.col - 1}, // NW
			{pos.row - 1, pos.col},     // N
			{pos.row - 1, pos.col + 1}, // NE
		},
		{
			{pos.row - 1, pos.col + 1}, // NE
			{pos.row, pos.col + 1},     // E
			{pos.row + 1, pos.col + 1}, // SE
		},
		{
			{pos.row + 1, pos.col + 1}, // SE
			{pos.row + 1, pos.col},     // S
			{pos.row + 1, pos.col - 1}, // SW
		},
		{
			{pos.row + 1, pos.col - 1}, // SW
			{pos.row, pos.col - 1},     // W
			{pos.row - 1, pos.col - 1}, // NW
		},
	}
}

func posWhenMoving(pos Pos, dir Direction) Pos {
	switch dir {
	case North:
		return Pos{pos.row - 1, pos.col}
	case East:
		return Pos{pos.row, pos.col + 1}
	case South:
		return Pos{pos.row + 1, pos.col}
	case West:
		return Pos{pos.row, pos.col - 1}
	}
	return pos
}

func directionsForRound(round uint) []Direction {
	round = round % 4
	return []Direction{North, South, West, East, North, South, West}[round : round+4]
}

func Print(elves PosMap[bool]) {
	padding := 1
	minRow, maxRow, minCol, maxCol := dimensions(elves)
	for row := minRow - padding; row <= maxRow+padding; row++ {
		for col := minCol - padding; col <= maxCol+padding; col++ {
			if elves[Pos{row, col}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	// fmt.Println()
}

func dimensions(elves PosMap[bool]) (minRow, maxRow, minCol, maxCol int) {
	minRow, maxRow, minCol, maxCol = math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for pos := range elves {
		minRow = utils.Min(minRow, pos.row)
		maxRow = utils.Max(maxRow, pos.row)
		minCol = utils.Min(minCol, pos.col)
		maxCol = utils.Max(maxCol, pos.col)
	}
	return
}

func score(elves PosMap[bool]) int {
	minRow, maxRow, minCol, maxCol := dimensions(elves)
	return (maxRow-minRow+1)*(maxCol-minCol+1) - len(elves)
}

func parseInput(input []byte) PosMap[bool] {
	result := make(PosMap[bool])
	for row, line := range strings.Split(string(input), "\n") {
		for col, value := range strings.Split(line, "") {
			if value == "#" {
				result[Pos{row, col}] = true
			}
		}
	}
	return result
}
