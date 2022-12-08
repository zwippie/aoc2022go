package myinput_test

import (
	"aoc2022/myinput"
	"testing"
)

func TestReadInput(t *testing.T) {
	expected := []byte{49, 10, 50, 10, 51}
	got := myinput.ReadInput("tests/test.txt")

	if len(got) != len(expected) {
		t.Fatalf("lengths dont't match, expected: %v, got %v", len(expected), len(got))
	}
	for i, b := range expected {
		if got[i] != b {
			t.Fatalf("invalid value, expected: %v, got: %v", expected, got)
		}
	}
}
