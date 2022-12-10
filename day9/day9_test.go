package day9_test

import (
	"aoc2022/day9"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day9.txt"
var expectedA = 6023
var expectedB = 2533

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day9.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day9.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day9.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day9.PartB(input)
	}
}
