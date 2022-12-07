package day6_test

import (
	"aoc2022/day6"
	"aoc2022/myinput"
	"testing"
)

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput("6", "")
	for i := 0; i < b.N; i++ {
		day6.PartA(input)
	}
}
