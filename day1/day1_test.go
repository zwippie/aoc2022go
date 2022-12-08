package day1_test

import (
	"aoc2022/day1"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day1.txt"
var expectedA = 71124
var expectedB = 204639

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day1.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day1.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day1.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day1.PartB(input)
	}
}
