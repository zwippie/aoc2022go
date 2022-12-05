package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []byte

type Move struct {
	count int
	from  int
	to    int
}

func PartA() {
	stacks, moves := readFile("input/day5.txt", 9)
	for _, move := range moves {
		stacks = doMove9000(stacks, move)
	}
	fmt.Println(getResult(stacks))
}

func PartB() {
	stacks, moves := readFile("input/day5.txt", 9)
	for _, move := range moves {
		stacks = doMove9001(stacks, move)
	}
	fmt.Println(getResult(stacks))
}

func doMove9000(stacks []Stack, move Move) []Stack {
	for i := 0; i < move.count; i++ {
		crate := stacks[move.from-1][len(stacks[move.from-1])-1]
		stacks[move.to-1] = append(stacks[move.to-1], crate)
		stacks[move.from-1] = stacks[move.from-1][:len(stacks[move.from-1])-1]
	}
	return stacks
}

func doMove9001(stacks []Stack, move Move) []Stack {
	crates := stacks[move.from-1][len(stacks[move.from-1])-move.count:]
	stacks[move.to-1] = append(stacks[move.to-1], crates...)
	stacks[move.from-1] = stacks[move.from-1][:len(stacks[move.from-1])-move.count]
	return stacks
}

func getResult(stacks []Stack) string {
	result := []byte{}
	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}
	return string(result)
}

func readFile(fileName string, stackCount int) (stacks []Stack, moves []Move) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	temp := []string{}

	for scanner.Scan() {
		var line = scanner.Text()
		if strings.HasPrefix(line, "move") {
			moves = append(moves, parseMove(line))
		} else if line != "" {
			temp = append(temp, line)
		}
	}
	stacks = parseStacks(temp, stackCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func parseMove(line string) Move {
	parts := strings.Split(line, " ")
	count, _ := strconv.Atoi(parts[1])
	from, _ := strconv.Atoi(parts[3])
	to, _ := strconv.Atoi(parts[5])
	return Move{count, from, to}
}

func parseStacks(lines []string, stackCount int) (stacks []Stack) {
	stacks = make([]Stack, stackCount)
	for i := len(lines) - 2; i >= 0; i-- { // reversed, skip last
		for j := 0; j < stackCount; j++ {
			crate := lines[i][4*j+1]
			if crate != 32 {
				stacks[j] = append(stacks[j], crate)
			}
		}
	}
	return
}
