package day14_test

import (
	"aoc2022/day14"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day14.txt"
var expectedA = 888
var expectedB = 26461

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day14.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day14.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day14.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day14.PartB(input)
	}
}
