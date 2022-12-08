package main

import (
	"aoc2022/day1"
	"aoc2022/day2"
	"aoc2022/day3"
	"aoc2022/day4"
	"aoc2022/day5"
	"aoc2022/day6"
	"aoc2022/day7"
	"aoc2022/day8"
	"aoc2022/myinput"
	"log"
	"os"
)

var days = map[string]func([]byte){
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
	"6a": day6.PartA,
	"6b": day6.PartB,
	"7a": day7.PartA,
	"7b": day7.PartB,
	"8a": day8.PartA,
	"8b": day8.PartB,
}

func main() {
	day := os.Args[1]
	part := os.Args[2]
	inputSuffix := ""
	if len(os.Args) > 3 {
		inputSuffix = os.Args[3] // day3example.txt
	}

	if f, ok := days[day+part]; ok {
		data := myinput.ReadInput(day, inputSuffix)
		f(data)
	} else {
		log.Fatal(day, " not implemented")
	}
}
