package day13

import (
	"aoc2022/stack"
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

type Element interface {
	int | List
}

type List []any

func PartA(input []byte) any {
	input = []byte("[[1],[2,3,4]]")
	data := parseInput(input)
	fmt.Printf("data: %#v\n", data)
	fmt.Printf("len(data): %v\n", len(data))
	fmt.Printf("data[0]: %T %v\n", data[0], data[0])
	fmt.Printf("len(data[0]): %v\n", len(data[0]))
	fmt.Printf("data[0][0]: %T %v\n", data[0][0], data[0][0])

	// len([int]List(data[0][0]))
	reflectOnList((data[0][0]))

	var s = []interface{}{1, 2, "three", reflectOnList}
	for k, v := range s {
		fmt.Printf("k: %v, v: %v, T(v): %T\n", k, v, v)
	}

	return 0
}

func PartB(input []byte) any {
	return 0
}

func reflectOnList(list any) {
	fmt.Printf("reflectOnList: %v\n", list)
	kind := reflect.TypeOf(list).Kind()
	fmt.Printf("kind: %v\n", kind)
	switch kind {
	case reflect.Slice:
		fmt.Println("it's a slice")
		s := reflect.ValueOf(list)
		fmt.Println(s, s.Len())

		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
			reflectOnList(s.Index(i))
		}
	case reflect.Int:
		fmt.Println("it's an int")
		s := reflect.ValueOf(list)
		fmt.Println(s)
	case reflect.Struct:
		fmt.Println("it's a struct")
		s := reflect.ValueOf(list)
		fmt.Println(s)
	case reflect.Array:
		fmt.Println("it's an array")
		s := reflect.ValueOf(list)
		fmt.Println(s)
	case reflect.Interface:
		fmt.Println("it's an interface")
		s := reflect.ValueOf(list)
		fmt.Println(s)
	}
}

func parseInput(input []byte) []List {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	result := []List{}

	for scanner.Scan() {
		line := scanner.Bytes()
		list := parseList(line)
		result = append(result, list)
		break
	}

	return result
}

func parseList(input []byte) List {
	listStack := stack.Stack[List]{}
	digits := []byte{}
	list := List{}

	for _, char := range input {
		fmt.Printf("char: %v, %s\n", char, string(char))

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
	fmt.Printf("listStack.IsEmpty(): %v\n", listStack.IsEmpty())
	return list
}
