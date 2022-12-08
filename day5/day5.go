package day5

import (
	"aoc2022/stack"
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
)

type Move struct {
	count int
	from  int
	to    int
}

// type Stack stack.Stack[byte]

func PartA(input []byte) any {
	stacks, moves := readFile(input, 9)
	for _, move := range moves {
		stacks = doMove9000(stacks, move)
	}
	return getResult(stacks) // WHTLRMZRC
}

func PartB(input []byte) any {
	stacks, moves := readFile(input, 9)
	for _, move := range moves {
		stacks = doMove9001(stacks, move)
	}
	return getResult(stacks) // GMPMLWNMG
}

func doMove9000(stacks []stack.Stack[byte], move Move) []stack.Stack[byte] {
	for i := 0; i < move.count; i++ {
		stacks[move.to].Push(stacks[move.from].Pop())
	}
	return stacks
}

func doMove9001(stacks []stack.Stack[byte], move Move) []stack.Stack[byte] {
	temp := make(stack.Stack[byte], move.count)
	for i := 0; i < move.count; i++ {
		temp.Push(stacks[move.from].Pop())
	}
	for i := 0; i < move.count; i++ {
		stacks[move.to].Push(temp.Pop())
	}
	return stacks
}

func getResult(stacks []stack.Stack[byte]) string {
	result := []byte{}
	for _, s := range stacks {
		result = append(result, s.Pop())
	}
	return string(result)
}

func readFile(input []byte, stackCount int) (stacks []stack.Stack[byte], moves []Move) {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
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
	return Move{count, from - 1, to - 1}
}

func parseStacks(lines []string, stackCount int) (stacks []stack.Stack[byte]) {
	stacks = make([]stack.Stack[byte], stackCount)
	for i := len(lines) - 2; i >= 0; i-- { // reversed, skip last
		for j := 0; j < stackCount; j++ {
			if crate := lines[i][4*j+1]; crate != 32 {
				stacks[j].Push(crate)
			}
		}
	}
	return
}
