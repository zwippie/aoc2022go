package day13

import (
	"aoc2022/stack"
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

type Elem interface {
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
	d := reflectOn((data[0][0]))
	fmt.Printf("d: %#v\n", d)

	return 0
}

func PartB(input []byte) any {
	return 0
}

func reflectOn(value any) any {
	switch v := value.(type) {
	case int:
		fmt.Printf("type of %v is %v int\n", value, v)
		return int(v)
	default:
		fmt.Printf("type of %v is %v so must be a List?\n", value, v)
		l := v.(List)
		for k, v := range l {
			fmt.Printf("k: %v, v: %v\n", k, v)
			fmt.Printf("%v: %#v\n", v, reflectOn(v))
		}
		return l
	}
}

func reflectOnList(value any) {
	fmt.Printf("reflectOnList: %v\n", value)
	kind := reflect.TypeOf(value).Kind()
	fmt.Printf("kind: %v\n", kind)
	switch kind {
	case reflect.Slice:
		fmt.Println("it's a slice")
		s := reflect.ValueOf(value)
		fmt.Println(s, s.Len())

		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
			reflectOnList(s.Index(i))
		}
	case reflect.Int:
		fmt.Println("it's an int")
		s := reflect.ValueOf(value)
		fmt.Println(s)
	case reflect.Struct:
		fmt.Println("it's a struct")
		s := reflect.ValueOf(value)
		fmt.Println(s)
		// list := List(value)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
			reflectOnList(s.Index(i))
		}
	case reflect.Array:
		fmt.Println("it's an array")
		s := reflect.ValueOf(value)
		fmt.Println(s)
	case reflect.Interface:
		fmt.Println("it's an interface")
		s := reflect.ValueOf(value)
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
