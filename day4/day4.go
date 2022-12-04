package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func PartA() {
	data := readFile("input/day4.txt")
	result := 0

	for _, pair := range data {
		if fullyContains(pair[0], pair[1]) || fullyContains(pair[1], pair[0]) {
			result += 1
		}
	}

	fmt.Println(result)
}

func PartB() {
	data := readFile("input/day4.txt")
	result := 0

	for _, pair := range data {
		if hasOverlap(pair[0], pair[1]) {
			result += 1
		}
	}

	fmt.Println(result)
}

// Does first contain all elements present in second?
func fullyContains(first []int, second []int) bool {
	for _, val := range second {
		if !slices.Contains(first, val) {
			return false
		}
	}
	return true
}

// Does first contain any element present in second?
func hasOverlap(first []int, second []int) bool {
	for _, val := range second {
		if slices.Contains(first, val) {
			return true
		}
	}
	return false
}

func readFile(fileName string) [][][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := [][][]int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = strings.Split(scanner.Text(), ",")
		result = append(result, [][]int{intRange(line[0]), intRange(line[1])})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func intRange(fromTo string) []int {
	firstLast := strings.Split(fromTo, "-")
	first, _ := strconv.Atoi(firstLast[0])
	last, _ := strconv.Atoi(firstLast[1])
	result := []int{}
	for i := first; i <= last; i++ {
		result = append(result, i)
	}
	return result
}
