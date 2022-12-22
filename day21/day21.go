package day21

// Directed acyclic graph (DAG)
// For part b, this was useful: https://www.cs.utexas.edu/users/novak/algebra.pdf

import (
	"fmt"
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
	Equal
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
	case Equal:
		return "="
	}
	return ""
}

type Node struct {
	name     string
	operator Operator
	left     *Node
	right    *Node
	value    int
	variable bool
}

func (node *Node) String() string {
	if node.variable {
		return "X"
	}
	if node.left == nil && node.right == nil {
		return fmt.Sprint(node.value)
	}
	return "(" + node.operator.String() + " " + node.left.String() + " " + node.right.String() + ")"
}

type NodeMap map[string]*Node

// 324122188240430
func PartA(input []byte) any {
	nodes := parseInput(input)
	return resolveValue(nodes["root"])
}

// 3412650897405
func PartB(input []byte) any {
	nodes := parseInput(input)

	root := nodes["root"]
	root.operator = Equal
	nodes["humn"].variable = true
	// fmt.Printf("equation: %v\n", nodes["root"])

	root.left, root.right = root.right, root.left
	solved := solveEquation(root)
	// fmt.Printf("solved: %v\n", solved)

	return resolveValue(solved.right)
}

func resolveValue(node *Node) int {
	switch node.operator {
	case Add:
		return resolveValue(node.left) + resolveValue(node.right)
	case Subtract:
		return resolveValue(node.left) - resolveValue(node.right)
	case Multiply:
		return resolveValue(node.left) * resolveValue(node.right)
	case Divide:
		return resolveValue(node.left) / resolveValue(node.right)
	}
	return node.value
}

func solveEquation(node *Node) *Node {
	// fmt.Printf("solving: %v\n", node)
	if node.left.variable {
		return node // found it on the left
	}
	if node.right.variable {
		node.left, node.right = node.right, node.left
		return node // found it, it was on the right
	}
	switch node.right.operator {
	case None:
		return nil // cannot be found with only value on the right side
	case Add:
		// y = a + b, try y - a = b and y - b = a
		copy := node.Copy()
		copy.left = &Node{left: copy.left, right: copy.right.left, operator: Subtract}
		copy.right = copy.right.right
		if copySolved := solveEquation(copy); copySolved != nil {
			return copySolved
		}
		copy = node.Copy()
		copy.left = &Node{left: copy.left, right: copy.right.right, operator: Subtract}
		copy.right = copy.right.left
		if copySolved := solveEquation(copy); copySolved != nil {
			return copySolved
		}
	case Subtract:
		// y = a - b, try y + b = a and  a - y = b
		copy := node.Copy()
		copy.left = &Node{left: copy.left, right: copy.right.right, operator: Add}
		copy.right = copy.right.left
		if copySolved := solveEquation(copy); copySolved != nil {
			return copySolved
		}
		copy = node.Copy()
		copy.left = &Node{left: copy.right.left, right: copy.left, operator: Subtract}
		copy.right = copy.right.right
		if copySolved := solveEquation(copy); copySolved != nil {
			return copySolved
		}
	case Multiply:
		// y = a * b, try y / a = b and y / b = a
		copy := node.Copy()
		copy.left = &Node{left: copy.left, right: copy.right.left, operator: Divide}
		copy.right = copy.right.right
		if copySolved := solveEquation(copy); copySolved != nil {
			return copySolved
		}
		copy = node.Copy()
		copy.left = &Node{left: copy.left, right: copy.right.right, operator: Divide}
		copy.right = copy.right.left
		if copySolved := solveEquation(copy); copySolved != nil {
			return copySolved
		}
	case Divide:
		// y = a / b, try y * b = a and a / y = b
		copy := node.Copy()
		copy.left = &Node{left: copy.left, right: copy.right.right, operator: Multiply}
		copy.right = copy.right.left
		if copySolved := solveEquation(copy); copySolved != nil {
			return copySolved
		}
		copy = node.Copy()
		copy.left = &Node{left: copy.right.left, right: copy.left, operator: Divide}
		copy.right = copy.right.right
		if copySolved := solveEquation(copy); copySolved != nil {
			return copySolved
		}
	}
	return nil // not found this way
}

func (node *Node) Copy() *Node {
	if node.operator == None {
		return &Node{
			name:     node.name,
			operator: node.operator,
			value:    node.value,
			variable: node.variable,
		}
	}
	return &Node{
		name:     node.name,
		operator: node.operator,
		left:     node.left.Copy(),
		right:    node.right.Copy(),
		value:    node.value,
		variable: node.variable,
	}
}

func parseInput(input []byte) NodeMap {
	nodes := make(NodeMap)
	for _, line := range strings.Split(string(input), "\n") {
		parts := strings.Split(line, ": ")
		name := parts[0]
		value, err := strconv.Atoi(parts[1])
		if err == nil {
			if nodes[name] == nil {
				nodes[name] = &Node{name: name, operator: None, value: value}
			} else {
				nodes[name].value = value
			}
		} else {
			formulaParts := strings.Split(parts[1], " ")
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
