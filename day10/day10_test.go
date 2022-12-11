package day10_test

import (
	"aoc2022/day10"
	"aoc2022/myinput"
	"testing"
)

var fileName = "day10.txt"
var expectedA = 17020
var expectedB = "###..#....####.####.####.#.....##..####.\n" +
	"#..#.#....#.......#.#....#....#..#.#....\n" +
	"#..#.#....###....#..###..#....#....###..\n" +
	"###..#....#.....#...#....#....#.##.#....\n" +
	"#.#..#....#....#....#....#....#..#.#....\n" +
	"#..#.####.####.####.#....####..###.####.\n"

func TestPartA(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day10.PartA(input)
	if got != expectedA {
		t.Errorf("expected: %v, got: %v", expectedA, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day10.PartA(input)
	}
}

func TestPartB(t *testing.T) {
	input := myinput.ReadInput(fileName)
	got := day10.PartB(input)
	if got != expectedB {
		t.Errorf("expected: %v, got: %v", expectedB, got)
	}
}

func BenchmarkPartB(b *testing.B) {
	input := myinput.ReadInput(fileName)
	for i := 0; i < b.N; i++ {
		day10.PartB(input)
	}
}
