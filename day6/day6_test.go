package day6_test

import (
	"aoc2022/day6"
	"testing"
)

func BenchmarkPartA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6.PartA()
	}
}
