package day12

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type Pos struct {
	row int
	col int
}
type Node struct {
	pos Pos
	val byte
}
type NodeMap map[Pos]*Node
type Path []Pos

var rows int
var cols int

// 425 is correct, but the current result 427 is too high...
func PartA(input []byte) any {
	nodes, startPos, endPos := parseInput(input)
	fmt.Printf("len(nodes): %v\n", len(nodes))
	fmt.Printf("nodes[start]: %#v\n", nodes[startPos])
	fmt.Printf("start: %v\n", startPos)
	fmt.Printf("end: %v\n", endPos)

	dist, prev := dijkstra(nodes, startPos, endPos)
	fmt.Printf("len(prev): %v\n", len(prev))
	fmt.Printf("dist[end]: %v\n", dist[endPos])

	path := getPath(prev, startPos, endPos, nodes)
	printMap(nodes, path)

	return dist[endPos] // the answer
}

func PartB(input []byte) any {
	return 0
}

// dijkstra shortest path algorithm, as taken from wikipedia
func dijkstra(nodes NodeMap, startPos Pos, endPos Pos) (map[Pos]int, map[*Node]*Node) {
	dist := make(map[Pos]int)     // current distance of each pos to startPos
	prev := make(map[*Node]*Node) // pointers to previous nodes on the shortest path
	q := make(NodeMap)            // map of all the nodes stil to check

	for pos, node := range nodes {
		dist[pos] = math.MaxInt // initial distance from start is very large
		prev[node] = nil        // all nodes have no prev set yet (can be skipped?)
		q[pos] = node           // add all nodes to the q
	}
	dist[startPos] = 0 // set distance of start pos to 0 so it will be picked first

	// if q should be empty then there is no path from start to end
	for len(q) > 0 {
		// what node in the q has the shortest distance to start?
		currentPos, ok := posWithMinDistance(q, dist)
		if !ok {
			fmt.Printf("len(q): %v\n", len(q))
			log.Fatal("could not find pos with min distance")
		}
		if currentPos == endPos {
			fmt.Println("found it!")
			return dist, prev
		}
		currentNode := q[currentPos]
		// remove the current node from the q
		// any node trying to visit this node later will always have a longer dist to start
		delete(q, currentPos)

		for _, pos := range neighboursStillInQ(q, currentNode) {
			if node, ok := q[pos]; ok {
				alt := dist[currentNode.pos] + 1
				if alt < dist[node.pos] {
					dist[node.pos] = alt
					prev[node] = currentNode
				}
			}
		}
	}
	return dist, prev
}

func posWithMinDistance(nodes NodeMap, dist map[Pos]int) (Pos, bool) {
	minDist, minPos, found := math.MaxInt, Pos{}, false
	for pos := range nodes {
		if dist[pos] < minDist {
			minDist = dist[pos]
			minPos = pos
			found = true
		}
	}
	return minPos, found
}

func neighboursStillInQ(q NodeMap, currentNode *Node) []Pos {
	result := []Pos{}
	val := currentNode.val

	if node, ok := q[Pos{currentNode.pos.row - 1, currentNode.pos.col}]; ok {
		if node.val <= val+1 {
			result = append(result, node.pos)
		}
	}
	if node, ok := q[Pos{currentNode.pos.row + 1, currentNode.pos.col}]; ok {
		if node.val <= val+1 {
			result = append(result, node.pos)
		}
	}
	if node, ok := q[Pos{currentNode.pos.row, currentNode.pos.col - 1}]; ok {
		if node.val <= val+1 {
			result = append(result, node.pos)
		}
	}
	if node, ok := q[Pos{currentNode.pos.row, currentNode.pos.col + 1}]; ok {
		if node.val <= val+1 {
			result = append(result, node.pos)
		}
	}
	return result
}

// recreate path from prev map
func getPath(prev map[*Node]*Node, startPos Pos, endPos Pos, nodes NodeMap) []Pos {
	result := []Pos{}
	u := nodes[endPos]
	if prev[u] != nil || u.pos == startPos {
		for u != nil {
			result = append([]Pos{u.pos}, result...)
			u = prev[u]
		}
	}
	return result
}

func printMap(nodes NodeMap, path []Pos) {
	pathMap := make(map[Pos]bool)
	for _, pos := range path {
		pathMap[pos] = true
	}

	result := ""
	for row := 0; row < rows; row++ {
		line := ""
		for col := 0; col < cols; col++ {
			char := nodes[Pos{row, col}].val
			if pathMap[Pos{row, col}] { // on path
				line += fmt.Sprintf("\033[0;%dm%s\033[m", 41, string(char))
			} else {
				if char == 99 { // hide sea/c
					line += " "
				} else if char == 98 { // wall/b
					line += fmt.Sprintf("\033[0;%dm%s\033[m", 44, string(char))
				} else {
					line += string(char)
				}
			}
		}
		result += line + "\n"
	}
	fmt.Println(result)
}

func parseInput(input []byte) (NodeMap, Pos, Pos) {
	grid := [][]byte{}
	startPos, endPos := Pos{}, Pos{}
	parts := strings.Split(string(input), "\n")
	for r, part := range parts {
		row := []byte{}
		for c, b := range []byte(part) {
			if b == 83 { // S
				startPos = Pos{r, c}
				b = 96 // one before a
			}
			if b == 69 { // E
				endPos = Pos{r, c}
				b = 123 // one after z
			}
			row = append(row, b)
		}
		grid = append(grid, row)
	}
	rows = len(grid)
	cols = len(grid[0])
	fmt.Printf("rows: %v, cols: %v\n", rows, cols)

	nodes := make(map[Pos]*Node)
	for row := range grid {
		for col, val := range grid[row] {
			pos := Pos{row, col}
			nodes[pos] = &Node{pos: pos, val: val}
		}
	}

	return nodes, startPos, endPos
}
