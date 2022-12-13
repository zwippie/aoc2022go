package day13

import (
	"aoc2022/stack"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
)

type List []any

func PartA(input []byte) any {
	pairs := parseInput(input)
	fmt.Printf("packets: %v\n", pairs)

	result := 0
	for i, pair := range pairs {
		fmt.Println("START of", i+1)
		left, right := pair[0], pair[1]
		_, inOrder := inRightOrder(left, right)
		if inOrder {
			result += i + 1
		}
	}

	return result
}

func PartB(input []byte) any {
	return 0
}

func inRightOrder(left List, right List) (cont bool, inOrder bool) {
	fmt.Printf("left: %v\n", left)
	fmt.Printf("right: %v\n", right)

	if left.IsEmpty() && right.IsEmpty() {
		fmt.Println("Both lists are empty")
		return true, true // continue
	}
	if left.IsEmpty() && !right.IsEmpty() {
		fmt.Println("Left is empty, right not")
		return false, true // right order
	}
	if !left.IsEmpty() && right.IsEmpty() {
		fmt.Println("Right is empty, left not")
		return false, false // right order
	}

	if left.NextIsInt() && right.NextIsInt() {
		fmt.Println("Both are integers")
		_, left, lVal, _ := left.Next()
		_, right, rVal, _ := right.Next()
		if lVal < rVal {
			return false, true
		} else if lVal > rVal {
			return false, false
		}
		fmt.Println("  But they are equal")
		return inRightOrder(left, right)
	}

	if left.NextIsInt() && !right.NextIsInt() {
		fmt.Println("Wrap the left")
		_, left, lVal, _ := left.Next()
		rNext, right, _, _ := right.Next()
		lNext := List{lVal}
		if cont, inOrder := inRightOrder(lNext, rNext); cont {
			return inRightOrder(left, right)
		} else {
			return cont, inOrder
		}
	}

	if !left.NextIsInt() && right.NextIsInt() {
		fmt.Println("Wrap the right")
		lNext, left, _, _ := left.Next()
		_, right, rVal, _ := right.Next()
		rNext := List{rVal}
		if cont, inOrder := inRightOrder(lNext, rNext); cont {
			return inRightOrder(left, right)
		} else {
			return cont, inOrder
		}
	}

	if !left.NextIsInt() && !right.NextIsInt() {
		fmt.Println("Both are lists")
		lNext, left, _, _ := left.Next()
		rNext, right, _, _ := right.Next()
		if cont, inOrder := inRightOrder(lNext, rNext); cont {
			return inRightOrder(left, right)
		} else {
			return cont, inOrder
		}
	}

	fmt.Println(">>End of comparison")
	return false, false // never reached?
}

func (l List) IsEmpty() bool {
	return len(l) == 0
}

func (l List) Next() (next List, rest List, num int, ok bool) {
	if l.IsEmpty() {
		return // List{}, 0, l, false
	}
	head, tail := l[0], l[1:]
	switch v := head.(type) {
	case int:
		return List{}, tail, int(v), true
	}
	return head.(List), tail, 0, true
}

func (l List) NextIsInt() bool {
	if l.IsEmpty() {
		return false
	}
	switch l[0].(type) {
	case int:
		return true
	}
	return false
}

func parseInput(input []byte) [][]List {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	result := [][]List{}

	for scanner.Scan() {
		pair := []List{}
		line := scanner.Bytes()
		list := parseList(line)
		pair = append(pair, list)
		scanner.Scan()
		line = scanner.Bytes()
		list = parseList(line)
		pair = append(pair, list)
		result = append(result, pair)
		scanner.Scan()
	}

	return result
}

func parseList(input []byte) List {
	listStack := stack.Stack[List]{}
	digits := []byte{}
	list := List{}

	if len(input) == 0 {
		return list
	}

	for _, char := range input { //[1 : len(input)-1] {
		// fmt.Printf("char: %v, %s\n", char, string(char))

		switch char {
		case 10: // newline
			return list
		case 44: // comma
			if len(digits) > 0 {
				number, _ := strconv.Atoi(string(digits))
				list = append(list, number)
				digits = []byte{}
			}
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57: // digit
			digits = append(digits, char)
		case 91: // [
			listStack.Push(list)
			list = List{}

		case 93: // ]
			if len(digits) > 0 {
				number, _ := strconv.Atoi(string(digits))
				list = append(list, number)
				digits = []byte{}
			}
			prevList := listStack.Pop()
			prevList = append(prevList, list)
			list = prevList
		}
	}
	if !listStack.IsEmpty() {
		log.Fatal("invalid input:", string(input))
	}
	return list[0].(List)
}
