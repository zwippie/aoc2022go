package day20

import (
	"aoc2022/utils"
	"container/ring"
	"fmt"
	"strings"
)

// not -18172
func PartA(input []byte) any {
	numbers, r := parseInput(input)

	for pos, num := range numbers {
		if num == 0 {
			continue // 0 does not move
		}
		atPos := moveToPos(r, pos)
		toMove := normalizeMove(num, len(numbers))
		otherPos := atPos.Move(toMove)
		// fmt.Println("move", numbers[atPos.Value.(int)], "between", numbers[otherPos.Value.(int)], "and", numbers[otherPos.Next().Value.(int)])
		if atPos != otherPos {
			otherPos.Link(atPos.Prev().Unlink(1))
		}
		r = otherPos
	}

	return getScore(r, numbers)
}

func PartB(input []byte) any {
	_, r := parseInput(input)

	r = moveToPos(r, 3127)
	fmt.Printf("r.Value: %v\n", r.Value)

	return 0
}

// move to the ring element that hold the pos value
func moveToPos(r *ring.Ring, pos int) *ring.Ring {
	for {
		if r.Value == pos {
			return r
		}
		r = r.Next()
	}
}

func normalizeMove(move int, length int) int {
	toMove := move
	if move < 0 {
		// blast
		toMove = ((move % length) + length - 1) % length
	}
	return toMove
}

func PrintRing(r *ring.Ring, numbers []int) {
	r.Do(func(pos any) {
		fmt.Print(numbers[pos.(int)], ", ")
	})
	fmt.Println()
}

func getScore(r *ring.Ring, original []int) int {
	result := 0
	zeroPos := 0
	for k, v := range original {
		if v == 0 {
			zeroPos = k
			break
		}
	}
	r = moveToPos(r, zeroPos)

	r = r.Move(1000)
	result += original[r.Value.(int)]
	fmt.Println(r.Value, original[r.Value.(int)])
	r = r.Move(1000)
	result += original[r.Value.(int)]
	fmt.Println(r.Value, original[r.Value.(int)])
	r = r.Move(1000)
	result += original[r.Value.(int)]
	fmt.Println(r.Value, original[r.Value.(int)])

	return result
}

func parseInput(input []byte) ([]int, *ring.Ring) {
	original := []int{}
	for _, line := range strings.Split(string(input), "\n") {
		number := utils.ParseInt(line)
		original = append(original, number)
	}
	r := ring.New(len(original))
	for idx := range original {
		r.Value = idx
		r = r.Next()
	}

	return original, r
}
