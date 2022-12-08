package day4_test

import (
	"aoc2022/day4"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day4.txt"
var expectedA = 532
var expectedB = 854

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day4.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day4.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day4.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day4.PartB(input)
	}
}
