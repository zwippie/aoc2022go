package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartA() {
	data := readFile("input/day2.txt")
	totalScore := 0
	for _, pair := range data {
		elf := normalizeInput(pair[0])
		player := normalizeInput(pair[1])
		totalScore += score(elf, player) + player + 1
	}
	fmt.Println(totalScore)
}

func PartB() {
	data := readFile("input/day2.txt")
	totalScore := 0
	for _, pair := range data {
		elf := normalizeInput(pair[0])
		hint := normalizeInput(pair[1])
		player := playerMove(elf, hint)
		totalScore += score(elf, player) + player + 1
	}
	fmt.Println(totalScore)
}

func playerMove(elf int, hint int) int {
	switch hint {
	case 0:
		return (elf + 2) % 3 // lose
	case 2:
		return (elf + 1) % 3 // win
	}
	return elf // draw
}

func score(elf int, player int) int {
	if elf == player {
		return 3
	}
	if (elf+1)%3 == player {
		return 6
	}
	return 0
}

func readFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := [][]string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = scanner.Text()
		pair := strings.Split(line, " ")
		result = append(result, pair)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func normalizeInput(value string) int {
	switch value {
	case "A", "X":
		return 0
	case "B", "Y":
		return 1
	case "C", "Z":
		return 2
	default:
		log.Fatal("invalid shape ", value)
	}
	return 0
}
