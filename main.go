package main

import (
	"aoc2022/day1"
	"aoc2022/day10"
	"aoc2022/day11"
	"aoc2022/day12"
	"aoc2022/day13"
	"aoc2022/day14"
	"aoc2022/day15"
	"aoc2022/day16"
	"aoc2022/day17"
	"aoc2022/day18"
	"aoc2022/day2"
	"aoc2022/day20"
	"aoc2022/day21"
	"aoc2022/day22"
	"aoc2022/day23"
	"aoc2022/day3"
	"aoc2022/day4"
	"aoc2022/day5"
	"aoc2022/day6"
	"aoc2022/day7"
	"aoc2022/day8"
	"aoc2022/day9"
	"aoc2022/myinput"
	"fmt"
	"log"
	"os"
)

var days = map[string]func([]byte) any{
	"1a":  day1.PartA,
	"1b":  day1.PartB,
	"2a":  day2.PartA,
	"2b":  day2.PartB,
	"3a":  day3.PartA,
	"3b":  day3.PartB,
	"4a":  day4.PartA,
	"4b":  day4.PartB,
	"5a":  day5.PartA,
	"5b":  day5.PartB,
	"6a":  day6.PartA,
	"6b":  day6.PartB,
	"7a":  day7.PartA,
	"7b":  day7.PartB,
	"8a":  day8.PartA,
	"8b":  day8.PartB,
	"9a":  day9.PartA,
	"9b":  day9.PartB,
	"10a": day10.PartA,
	"10b": day10.PartB,
	"11a": day11.PartA,
	"11b": day11.PartB,
	"12a": day12.PartA,
	"12b": day12.PartB,
	"13a": day13.PartA,
	"13b": day13.PartB,
	"14a": day14.PartA,
	"14b": day14.PartB,
	"15a": day15.PartA,
	"15b": day15.PartB,
	"16a": day16.PartA,
	"16b": day16.PartB,
	"17a": day17.PartA,
	"17b": day17.PartB,
	"18a": day18.PartA,
	"18b": day18.PartB,
	"20a": day20.PartA,
	"20b": day20.PartB,
	"21a": day21.PartA,
	"21b": day21.PartB,
	"22a": day22.PartA,
	"22b": day22.PartB,
	"23a": day23.PartA,
	"23b": day23.PartB,
}

func main() {
	day := os.Args[1]
	part := os.Args[2]
	inputSuffix := ""
	if len(os.Args) > 3 {
		inputSuffix = os.Args[3] // day3example.txt
	}

	if f, ok := days[day+part]; ok {
		fileName := fmt.Sprintf("day%s%s.txt", day, inputSuffix)
		data := myinput.ReadInput(fileName)
		fmt.Println(f(data))
	} else {
		log.Fatalf("day %s%s not implemented", day, part)
	}
}
