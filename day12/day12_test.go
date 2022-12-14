package day12_test

import (
	"aoc2022/day12"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day12.txt"
var expectedA = 427
var expectedB = 0

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day12.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day12.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day12.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day12.PartB(input)
	}
}
