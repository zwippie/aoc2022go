package day1_test

import (
	"aoc2022/day1"
	"aoc2022/myinput"
	"testing"
)

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput("1", "")
	for i := 0; i < b.N; i++ {
		day1.PartA(input)
	}
}

func TestPartA(t *testing.T) {
	input := myinput.ReadInput("1", "")
	day1.PartA(input) // must return to assert...
}
