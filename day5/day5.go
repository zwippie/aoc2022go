package day5

import (
	"aoc2022/stack"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	count int
	from  int
	to    int
}

// type Stack stack.Stack[byte]

func PartA() {
	stacks, moves := readFile("input/day5.txt", 9)
	for _, move := range moves {
		stacks = doMove9000(stacks, move)
	}
	printResult(stacks) // WHTLRMZRC
}

func PartB() {
	stacks, moves := readFile("input/day5.txt", 9)
	for _, move := range moves {
		stacks = doMove9001(stacks, move)
	}
	printResult(stacks) // GMPMLWNMG
}

func doMove9000(stacks []stack.Stack[byte], move Move) []stack.Stack[byte] {
	for i := 0; i < move.count; i++ {
		value := stacks[move.from-1].Pop()
		stacks[move.to-1].Push(value)
	}
	return stacks
}

func doMove9001(stacks []stack.Stack[byte], move Move) []stack.Stack[byte] {
	temp := make(stack.Stack[byte], move.count)

	for i := 0; i < move.count; i++ {
		value := stacks[move.from-1].Pop()
		temp.Push(value)
	}
	for i := 0; i < move.count; i++ {
		value := temp.Pop()
		stacks[move.to-1].Push(value)
	}

	return stacks
}

func printResult(stacks []stack.Stack[byte]) {
	result := []byte{}
	for _, s := range stacks {
		result = append(result, s.Pop())
	}
	fmt.Println(string(result))
}

func readFile(fileName string, stackCount int) (stacks []stack.Stack[byte], moves []Move) {
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

func parseStacks(lines []string, stackCount int) (stacks []stack.Stack[byte]) {
	stacks = make([]stack.Stack[byte], stackCount)
	for i := len(lines) - 2; i >= 0; i-- { // reversed, skip last
		for j := 0; j < stackCount; j++ {
			crate := lines[i][4*j+1]
			if crate != 32 {
				stacks[j].Push(crate)
			}
		}
	}
	return
}
