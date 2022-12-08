package day8_test

import (
	"aoc2022/day8"
	"aoc2022/myinput"
	"testing"
)

var expectedA = 1700
var expectedB = 470596

func TestPartA(t *testing.T) {
	input := myinput.ReadInput("8", "")
	got := day8.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput("8", "")
	for i := 0; i < b.N; i++ {
		day8.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput("8", "")
	got := day8.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput("8", "")
	for i := 0; i < b.N; i++ {
		day8.PartB(input)
	}
}
