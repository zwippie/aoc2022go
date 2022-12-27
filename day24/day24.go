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

type Blizzard struct {
	pos       Pos
	direction Direction
}

type Blizzards []*Blizzard // the blizzard input, used to build blizzardmaps

type BlizzardMap map[Pos]bool // does pos has one or more blizzards

type BlizzardTimeMap []BlizzardMap // the blizzard map for every minute (it's repeating)

var maxRow, maxCol, repeatingAt int
var entry, exit Pos
var btm BlizzardTimeMap
var minMinutes int = math.MaxInt

// 9223372036854775807 too high
func PartA(input []byte) any {
	btm = parseInput(input)
	findPath(1, entry)
	return minMinutes
}

func PartB(input []byte) any {
	return math.MaxInt
}

func findPath(minute int, player Pos) {
	if minute >= minMinutes {
		return
	}
	// if minute > 3*repeatingAt {
	// 	return
	// }
	if player == exit {
		minMinutes = utils.Min(minMinutes, minute-1)
		fmt.Printf("minMinutes: %v\n", minMinutes)
		return
	}
	blizzards := btm[minute%len(btm)]

	for _, p := range adjacent(player) {
		if !blizzards[p] {
			// fmt.Println(blizzards.ToString(p), minute)
			findPath(minute+1, p)
		}
	}
	if player != entry && !blizzards[player] { // always try to wait
		findPath(minute+1, player)
	}
}

func adjacent(pos Pos) []Pos {
	result := []Pos{}
	possible := []Pos{
		{pos.row - 1, pos.col},
		{pos.row, pos.col + 1},
		{pos.row + 1, pos.col},
		{pos.row, pos.col - 1},
	}
	for _, p := range possible {
		if (p.row >= 0 && p.row <= maxRow && p.col >= 0 && p.col <= maxCol) || p == exit {
			result = append(result, p)
		}
	}
	return result
}

func blizzardsAtMinute(blizzards Blizzards, minute int) BlizzardMap {
	result := make(BlizzardMap)
	rowLen, colLen := maxRow+1, maxCol+1
	for _, blizzard := range blizzards {
		switch blizzard.direction {
		case North:
			row := ((blizzard.pos.row-minute)%rowLen + rowLen) % rowLen
			result[Pos{row, blizzard.pos.col}] = true
		case East:
			col := (blizzard.pos.col + minute) % colLen
			result[Pos{blizzard.pos.row, col}] = true
		case South:
			row := (blizzard.pos.row + minute) % rowLen
			result[Pos{row, blizzard.pos.col}] = true
		case West:
			col := ((blizzard.pos.col-minute)%colLen + colLen) % colLen
			result[Pos{blizzard.pos.row, col}] = true
		}
	}
	return result
}

func (bm BlizzardMap) String() string {
	result := ""
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if bm[Pos{row, col}] {
				result += "X"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}

func (bm BlizzardMap) ToString(player Pos) string {
	result := ""
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if row == player.row && col == player.col {
				result += "O"
			} else if bm[Pos{row, col}] {
				result += "X"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}

func (d Direction) String() string {
	return string("^>v<"[d])
}

func parseInput(input []byte) BlizzardTimeMap {
	blizzards := Blizzards{}
	maxRow, maxCol = math.MinInt, math.MinInt
	for row, line := range strings.Split(string(input), "\n") {
		maxRow = utils.Max(maxRow, row)
		for col, value := range strings.Split(line, "") {
			maxCol = utils.Max(maxCol, col)
			if value == "^" {
				blizzard := &Blizzard{pos: Pos{row - 1, col - 1}, direction: North}
				blizzards = append(blizzards, blizzard)
			} else if value == ">" {
				blizzard := &Blizzard{pos: Pos{row - 1, col - 1}, direction: East}
				blizzards = append(blizzards, blizzard)
			} else if value == "v" {
				blizzard := &Blizzard{pos: Pos{row - 1, col - 1}, direction: South}
				blizzards = append(blizzards, blizzard)
			} else if value == "<" {
				blizzard := &Blizzard{pos: Pos{row - 1, col - 1}, direction: West}
				blizzards = append(blizzards, blizzard)
			}
		}
	}

	maxRow -= 2
	maxCol -= 2
	entry = Pos{-1, 0}
	exit = Pos{maxRow + 1, maxCol}

	repeatingAt = utils.LCM(maxRow+1, maxCol+1)
	btm := make(BlizzardTimeMap, repeatingAt)
	for minute := 0; minute < repeatingAt; minute++ {
		btm[minute] = blizzardsAtMinute(blizzards, minute)
	}
	return btm
}
