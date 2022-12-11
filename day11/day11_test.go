package day11_test

import (
	"aoc2022/day11"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day11.txt"
var expectedA = 117640
var expectedB = 30616425600

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day11.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day11.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day11.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day11.PartB(input)
	}
}
