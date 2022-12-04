package collection_test

import (
	"aoc2022/collection"
	"log"
	"testing"
)

func TestSumWithInts(t *testing.T) {
	data := []int{1, 2, 3}
	if collection.Sum(data) != 6 {
		log.Fatal("test error")
	}
}

func TestSumWithFloats(t *testing.T) {
	data := []float64{1.5, 2.5, 3.5}
	if collection.Sum(data) != 7.5 {
		log.Fatal("test error")
	}
}

func TestLastWithInts(t *testing.T) {
	data := []int{1, 2, 3}
	if collection.Last(data) != 3 {
		log.Fatal("test error")
	}
}

func TestLastWithRunes(t *testing.T) {
	data := []rune{97, 98, 99}
	if collection.Last(data) != 99 {
		log.Fatal("test error")
	}
}

func TestSliceMap(t *testing.T) {
	data := []int{1, 2, 3}
	double := func(x int) int {
		return 2 * x
	}
	collection.SliceMap(data, double)
}
