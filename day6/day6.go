package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PartA() {
	fmt.Println(findMarkerIdx(readFile("input/day6.txt"), 4))
}

func PartB() {
	fmt.Println(findMarkerIdx(readFile("input/day6.txt"), 14))
}

func findMarkerIdx(data []byte, markerSize int) int {
	lastIdx := map[byte]int{}
	length, start, end := 0, 0, 0
	for {
		start = max(start, lastIdx[data[end]]+1)
		length = max(length, end-start+1)
		lastIdx[data[end]] = end
		if length == markerSize {
			break
		}
		end++
	}
	return end + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func readFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	data, _ := reader.ReadBytes('\n')
	return data
}
