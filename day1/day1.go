package day1

import (
	"aoc2022/collection"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"sort"
	"strconv"
)

// 71124
func PartA(input []byte) {
	calories := readFile(input)

	sort.IntSlice(calories).Sort()
	fmt.Println(collection.Last(calories))
	// return collection.Last(calories)
}

// 204639
func PartB(input []byte) {
	calories := readFile(input)
	sort.IntSlice(calories).Sort()

	lastThree := collection.LastN(calories, 3)
	total := collection.Sum(lastThree)
	fmt.Println(total)
}

func readFile(input []byte) []int {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	result := []int{}
	calories := 0

	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			result = append(result, calories)
			calories = 0
		} else {
			number, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			calories += number
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
