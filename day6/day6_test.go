package day6_test

import (
	"aoc2022/day6"
	"aoc2022/myinput"
	"testing"
)

var expectedA = 1909
var expectedB = 3380

func TestPartA(t *testing.T) {
	input := myinput.ReadInput("6", "")
	got := day6.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput("6", "")
	for i := 0; i < b.N; i++ {
		day6.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput("6", "")
	got := day6.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput("6", "")
	for i := 0; i < b.N; i++ {
		day6.PartB(input)
	}
}
