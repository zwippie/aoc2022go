package day2_test

import (
	"aoc2022/day2"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day2.txt"
var expectedA = 13446
var expectedB = 13509

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day2.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day2.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day2.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day2.PartB(input)
	}
}
