package main

import (
	"aoc2022/day1"
	"aoc2022/day2"
	"aoc2022/day3"
	"aoc2022/day4"
	"aoc2022/day5"
	"log"
	"os"
)

var days = map[string]func(){
	"1a": day1.PartA,
	"1b": day1.PartB,
	"2a": day2.PartA,
	"2b": day2.PartB,
	"3a": day3.PartA,
	"3b": day3.PartB,
	"4a": day4.PartA,
	"4b": day4.PartB,
	"5a": day5.PartA,
	"5b": day5.PartB,
}

func main() {
	day := os.Args[1]

	if f, ok := days[day]; ok {
		f()
	} else {
		log.Fatal(day, " not implemented")
	}
}
