package day22

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

const (
	Right int = 0
	Down  int = 1
	Left  int = 2
	Up    int = 3
)

type Instruction struct {
	isTurn bool
	value  int
}

type Instructions []Instruction

type Pos struct {
	row, col int
}

type Node struct {
	pos                   Pos
	up, down, left, right *Node
	wall                  bool
}

type Maze map[Pos]*Node

type Player struct {
	node   *Node
	facing int
}

// 149250
func PartA(input []byte) any {
	_, entry, instructions := parseInput(input)
	player := Player{node: entry, facing: Right}
	for _, instruction := range instructions {
		player.PerformInstruction(instruction)
	}
	return player.Score()
}

func PartB(input []byte) any {
	return 0
}

func (player *Player) PerformInstruction(instruction Instruction) {
	if instruction.isTurn {
		player.Turn(instruction.value)
	} else {
		for i := 0; i < instruction.value; i++ {
			switch player.facing {
			case Right:
				if player.node.right.wall {
					break // cannot go any further
				}
				player.node = player.node.right
			case Down:
				if player.node.down.wall {
					break
				}
				player.node = player.node.down
			case Left:
				if player.node.left.wall {
					break
				}
				player.node = player.node.left
			case Up:
				if player.node.up.wall {
					break
				}
				player.node = player.node.up
			}
		}
	}
}

func (player *Player) Turn(direction int) {
	if direction == Right {
		player.facing = (player.facing + 1) % 4
	} else {
		player.facing = (player.facing + 3) % 4
	}
}

func (player *Player) Score() int {
	fmt.Println(player.node.pos.row, player.node.pos.col, player.facing)
	return (player.node.pos.row+1)*1000 + (player.node.pos.col+1)*4 + player.facing
}

func parseInput(input []byte) (maze Maze, entry *Node, instructions Instructions) {
	parts := strings.Split(string(input), "\n\n")
	maze, entry = parseMaze(parts[0])
	instructions = parseInstructions(parts[1])
	return
}

func parseMaze(input string) (maze Maze, entry *Node) {
	maze = make(map[Pos]*Node)

	for row, line := range strings.Split(input, "\n") {
		for col, val := range strings.Split(line, "") {
			if val != " " {
				node := &Node{pos: Pos{row, col}, wall: val == "#"}
				if entry == nil {
					entry = node
				}
				maze[Pos{row, col}] = node
			}
		}
	}
	maze = connectNodesA(maze)
	return
}

func connectNodesA(maze Maze) Maze {
	for pos, node := range maze {
		if node.up == nil {
			if n, ok := maze[Pos{pos.row - 1, pos.col}]; ok {
				node.up = n
				n.down = node
			} else {
				// find max row in this col
				maxRow := math.MinInt
				for p := range maze {
					if p.col == pos.col {
						maxRow = utils.Max(maxRow, p.row)
					}
				}
				node.up = maze[Pos{maxRow, pos.col}]
				maze[Pos{maxRow, pos.col}].down = node
			}
		}
		if node.left == nil {
			if n, ok := maze[Pos{pos.row, pos.col - 1}]; ok {
				node.left = n
				n.right = node
			} else {
				// find max col in this row
				maxCol := math.MinInt
				for p := range maze {
					if p.row == pos.row {
						maxCol = utils.Max(maxCol, p.col)
					}
				}
				node.left = maze[Pos{pos.row, maxCol}]
				maze[Pos{pos.row, maxCol}].right = node
			}
		}
	}

	return maze
}

func parseInstructions(input string) Instructions {
	values := regexp.MustCompile(`(\d+|L|R)?`).FindAllString(input, -1)
	result := Instructions{}
	for _, value := range values {
		if value == "L" {
			result = append(result, Instruction{isTurn: true, value: Left})
		} else if value == "R" {
			result = append(result, Instruction{isTurn: true, value: Right})
		} else {
			result = append(result, Instruction{value: utils.ParseInt(value)})
		}
	}
	return result
}
