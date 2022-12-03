package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartA() {
	totalScore := 0

	for _, line := range readFile("input/day3.txt") {
		runes := []rune(line)
		halfway := len(runes) / 2
		duplicate := findDuplicate(runes[:halfway], runes[halfway:])
		totalScore += priority(duplicate)
	}

	fmt.Println(totalScore)
}

func PartB() {
	data := readFile("input/day3.txt")
	totalScore := 0

	for i := 0; i < len(data); i += 3 {
		duplicate := findDuplicate3(data[i], data[i+1], data[i+2])
		totalScore += priority(duplicate)
	}

	fmt.Println(totalScore)
}

func priority(char rune) int {
	if char < 97 {
		return int(char) - 38
	}
	return int(char) - 96
}

func findDuplicate3(first string, second string, third string) rune {
	for _, char := range []rune(first) {
		if strings.ContainsRune(second, char) {
			if strings.ContainsRune(third, char) {
				return char
			}
		}
	}
	log.Fatal("no duplicate found")
	return 0
}

func findDuplicate(first []rune, second []rune) rune {
	for _, char := range first {
		if strings.ContainsRune(string(second), char) {
			return char
		}
	}
	log.Fatal("no duplicate found")
	return 0
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
