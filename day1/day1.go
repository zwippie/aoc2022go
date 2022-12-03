package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func PartA() {
	calories := readFile("input/day1.txt")

	sort.IntSlice(calories).Sort()
	fmt.Println(calories[len(calories)-1])

	// fmt.Println(collection.Collection(calories).Last())
	// collection.Test()
}

func PartB() {
	calories := readFile("input/day1.txt")
	sort.IntSlice(calories).Sort()

	lastThree := calories[len(calories)-3:]
	total := 0
	for _, value := range lastThree {
		total += value
	}
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
