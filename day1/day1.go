package day1

import (
	"aoc2022/collection"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// 71124
func PartA() {
	calories := readFile("input/day1.txt")

	sort.IntSlice(calories).Sort()
	fmt.Println(collection.Last(calories))
}

// 204639
func PartB() {
	calories := readFile("input/day1.txt")
	sort.IntSlice(calories).Sort()

	lastThree := collection.LastN(calories, 3)
	total := collection.Sum(lastThree)
	fmt.Println(total)
}

func readFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := []int{}
	calories := 0
	scanner := bufio.NewScanner(file)

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
