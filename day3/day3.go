package day3

import (
	"bufio"
	"bytes"
	"log"
	"strings"
)

// 7581
func PartA(input []byte) any {
	totalScore := 0

	for _, line := range readFile(input) {
		runes := []rune(line)
		halfway := len(runes) / 2
		duplicate := findDuplicate(runes[:halfway], runes[halfway:])
		totalScore += priority(duplicate)
	}

	return totalScore
}

// 2525
func PartB(input []byte) any {
	data := readFile(input)
	totalScore := 0

	for i := 0; i < len(data); i += 3 {
		duplicate := findDuplicate3(data[i], data[i+1], data[i+2])
		totalScore += priority(duplicate)
	}

	return totalScore
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

func readFile(input []byte) []string {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	result := []string{}

	for scanner.Scan() {
		var line = scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
