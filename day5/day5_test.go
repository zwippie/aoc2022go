package day5_test

import (
	"aoc2022/day5"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day5.txt"
var expectedA = "WHTLRMZRC"
var expectedB = "GMPMLWNMG"

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day5.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day5.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day5.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day5.PartB(input)
	}
}
