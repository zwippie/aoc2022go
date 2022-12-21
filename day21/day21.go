package day21

// Directed acyclic graph (DAG)
// And an RPN equation solver

import (
	"aoc2022/stack"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Operator int

const (
	None Operator = iota
	Add
	Subtract
	Multiply
	Divide
)

func (o Operator) String() string {
	switch o {
	case Add:
		return "+"
	case Subtract:
		return "-"
	case Multiply:
		return "*"
	case Divide:
		return "/"
	}
	return ""
}

type Node struct {
	name     string
	operator Operator
	left     *Node
	right    *Node
	resolved bool
	value    int
}

type NodeMap map[string]*Node

// 324122188240430
func PartA(input []byte) any {
	nodes := parseInput(input)
	return performOperation(nodes["root"])
}

func PartB(input []byte) any {
	nodes := parseInput(input)

	rpn := &[]any{}
	toRpn(nodes["root"], rpn)
	fmt.Printf("rpn: %v\n", rpn)
	result := evalRpn(rpn)
	fmt.Printf("result: %v\n", result)

	return 0 // smarterNaiveSearch(nodes)
}

func evalRpn(rpn *[]any) int {
	s := stack.Stack[int]{}
	for _, next := range *rpn {
		switch val := next.(type) {
		case int:
			s.Push(int(val))
		default:
			switch next {
			case Add:
				s.Push(s.Pop() + s.Pop())
			case Subtract:
				b, a := s.Pop(), s.Pop()
				s.Push(a - b)
			case Multiply:
				s.Push(s.Pop() * s.Pop())
			case Divide:
				b, a := s.Pop(), s.Pop()
				s.Push(a / b)
			}
		}
	}
	return int(s.Pop())
}

func toRpn(node *Node, result *[]any) {
	if node.left == nil && node.right == nil {
		*result = append(*result, node.value)
	} else {
		toRpn(node.left, result)
		toRpn(node.right, result)
		*result = append(*result, node.operator)
	}
}

func smarterNaiveSearch(nodes NodeMap) int {
	findLeft, findRight := findNode(nodes["root"].left, "humn"), findNode(nodes["root"].right, "humn")
	fmt.Printf("findLeft: %v\n", findLeft)
	fmt.Printf("findRight: %v\n", findRight)
	if findLeft == nil {
		leftNode := nodes["root"].left
		leftValue := performOperation(leftNode)
		newNode := &Node{name: leftNode.name, value: leftValue}
		nodes[leftNode.name] = newNode
		nodes["root"].left = newNode
	} else {
		rightNode := nodes["root"].right
		rightValue := performOperation(rightNode)
		newNode := &Node{name: rightNode.name, value: rightValue}
		nodes[rightNode.name] = newNode
		nodes["root"].right = newNode
	}

	return naiveSearch(nodes)
}

func naiveSearch(nodes NodeMap) int {
	nodes["root"].operator = Subtract // should return zero
	i := 0
	for {
		if i%1000000 == 0 {
			fmt.Printf("i: %v\n", i)
		}
		nodes["humn"].value = i
		if performOperation(nodes["root"]) == 0 {
			break
		}
		i++
	}
	return i
}

func performOperation(node *Node) int {
	switch node.operator {
	case Add:
		return performOperation(node.left) + performOperation(node.right)
	case Subtract:
		return performOperation(node.left) - performOperation(node.right)
	case Multiply:
		return performOperation(node.left) * performOperation(node.right)
	case Divide:
		return performOperation(node.left) / performOperation(node.right)
	}
	return node.value
}

func findNode(node *Node, name string) *Node {
	if node.name == name {
		return node
	}
	if node.left != nil {
		if found := findNode(node.left, name); found != nil {
			return found
		}
	}
	if node.right != nil {
		return findNode(node.right, name)
	}
	return nil // not found
}

func parseInput(input []byte) NodeMap {
	nodes := make(NodeMap)
	for _, line := range strings.Split(string(input), "\n") {
		parts := strings.Split(line, ": ")
		// fmt.Printf("parts: %#v\n", parts)
		name := parts[0]
		value, err := strconv.Atoi(parts[1])
		if err == nil {
			if nodes[name] == nil {
				nodes[name] = &Node{name: name, operator: None, value: value, resolved: true}
			} else {
				nodes[name].value = value
				nodes[name].resolved = true
			}
		} else {
			formulaParts := strings.Split(parts[1], " ")
			// fmt.Printf("formulaParts: %v\n", formulaParts)
			operator := Add
			switch formulaParts[1] {
			case "+":
				operator = Add
			case "-":
				operator = Subtract
			case "*":
				operator = Multiply
			case "/":
				operator = Divide
			default:
				log.Fatal("Unknown operator", formulaParts[1])
			}
			newNode := &Node{name: name, operator: operator}
			if nodes[name] != nil {
				newNode = nodes[name]
				newNode.operator = operator
			}
			leftName, rightName := formulaParts[0], formulaParts[2]

			if nodes[leftName] == nil {
				leftNode := &Node{name: leftName}
				nodes[leftName] = leftNode
				newNode.left = leftNode
			} else {
				newNode.left = nodes[leftName]
			}
			if nodes[rightName] == nil {
				rightNode := &Node{name: rightName}
				nodes[rightName] = rightNode
				newNode.right = rightNode
			} else {
				newNode.right = nodes[rightName]
			}
			nodes[name] = newNode
		}
	}

	return nodes
}
