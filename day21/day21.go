package day21

// Directed acyclic graph (DAG)?

import (
	"log"
	"strconv"
	"strings"
)

type Operator uint8

const (
	None Operator = iota
	Add
	Subtract
	Multiply
	Divide
)

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
	return 0
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
