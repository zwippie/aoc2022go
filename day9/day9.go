package day9

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

// 6023
func PartA(input []byte) any {
	data := parseInput(input)
	head, tail := Pos{0, 0}, Pos{0, 0}
	tailPath := map[Pos]int{}

	for _, dir := range data {
		head.move(dir)
		if !touching(head, tail) {
			tail.moveTail(head)
		}
		tailPath[tail] = 1
	}

	return len(tailPath)
}

// 2533
func PartB(input []byte) any {
	data := parseInput(input)
	var knots []Pos = make([]Pos, 10)
	tailPath := map[Pos]int{}

	for _, dir := range data {
		knots[0].move(dir)
		for idx := range knots {
			if idx != 0 {
				if !touching(knots[idx-1], knots[idx]) {
					knots[idx].moveTail(knots[idx-1])
				}
			}
		}
		tailPath[knots[9]] = 1
	}

	return len(tailPath)
}

func (knot *Pos) move(dir string) {
	switch dir {
	case "U":
		knot.y -= 1
	case "R":
		knot.x += 1
	case "D":
		knot.y += 1
	case "L":
		knot.x -= 1
	}
}

func touching(a Pos, b Pos) bool {
	return abs(a.x-b.x) <= 1 && abs(a.y-b.y) <= 1
}

func (tail *Pos) moveTail(head Pos) {
	dx := head.x - tail.x
	dy := head.y - tail.y
	if dx > 0 {
		tail.x += 1
	} else if dx < 0 {
		tail.x -= 1
	}
	if dy > 0 {
		tail.y += 1
	} else if dy < 0 {
		tail.y -= 1
	}
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func parseInput(input []byte) []string {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	result := []string{}

	for scanner.Scan() {
		var line = scanner.Text()
		parts := strings.Split(line, " ")
		steps, _ := strconv.Atoi(parts[1])
		for s := 0; s < steps; s++ {
			result = append(result, parts[0])
		}
	}

	return result
}
