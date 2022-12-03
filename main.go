package main

import (
	"aoc2022/day1"
	"aoc2022/day2"
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]
	fmt.Println("It's day", day)

	switch day {
	case "1a":
		day1.PartA()
	case "1b":
		day1.PartB()
	case "2a":
		day2.PartA()
	case "2b":
		day2.PartB()
	default:
		fmt.Println(day, "not found")
	}
}
