package day5_test

import (
	"aoc2022/day5"
	"aoc2022/myinput"
	"testing"
)

var expectedA = "WHTLRMZRC"
var expectedB = "GMPMLWNMG"

func TestPartA(t *testing.T) {
	input := myinput.ReadInput("5", "")
	got := day5.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput("5", "")
	for i := 0; i < b.N; i++ {
		day5.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput("5", "")
	got := day5.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput("5", "")
	for i := 0; i < b.N; i++ {
		day5.PartB(input)
	}
}
