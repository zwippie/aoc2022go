package day13_test

import (
	"aoc2022/day13"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day13.txt"
var expectedA = 5252
var expectedB = 20592

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day13.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day13.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day13.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day13.PartB(input)
	}
}
