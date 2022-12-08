package day3_test

import (
	"aoc2022/day3"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day3.txt"
var expectedA = 7581
var expectedB = 2525

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day3.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day3.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day3.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day3.PartB(input)
	}
}
