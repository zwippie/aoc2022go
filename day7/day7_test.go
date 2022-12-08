package day7_test

import (
	"aoc2022/day7"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day7.txt"
var expectedA = 1611443
var expectedB = 2086088

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day7.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day7.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day7.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day7.PartB(input)
	}
}
