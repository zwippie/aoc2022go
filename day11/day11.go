package day11

import (
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id          int
	items       []int
	operation   func(int) int
	divisibleBy int
	ifTrue      int
	ifFalse     int
	inspected   int
}

type Monkeys []*Monkey

// 117640
func PartA(input []byte) any {
	monkeys, dividerMul := parseInput(input)
	for i := 0; i < 20; i++ {
		monkeys.performRound(true, dividerMul)
	}
	return getScore(monkeys)
}

// 30616425600
func PartB(input []byte) any {
	monkeys, dividerMul := parseInput(input)
	for i := 0; i < 10000; i++ {
		monkeys.performRound(false, dividerMul)
	}
	return getScore(monkeys)
}

func (monkeys Monkeys) performRound(divideByThree bool, dividerMul int) {
	for _, monkey := range monkeys {
		for _, item := range monkey.items {
			item = monkey.operation(item)
			monkey.inspected++
			if divideByThree {
				item /= 3
			}
			item %= dividerMul
			if item%monkey.divisibleBy == 0 {
				otherMonkey := monkeys[monkey.ifTrue]
				otherMonkey.items = append(otherMonkey.items, item)
			} else {
				otherMonkey := monkeys[monkey.ifFalse]
				otherMonkey.items = append(otherMonkey.items, item)
			}
		}
		monkey.items = []int{}
	}
}

func getScore(monkeys Monkeys) int {
	inspected := []int{}
	for _, monkey := range monkeys {
		inspected = append(inspected, monkey.inspected)
	}
	sort.Ints(inspected)
	return inspected[len(inspected)-2] * inspected[len(inspected)-1]
}

func parseInput(input []byte) (Monkeys, int) {
	monkeys := Monkeys{}
	parts := strings.Split(string(input), "\n\n")

	for _, part := range parts {
		monkeys = append(monkeys, parseMonkey(part))
	}

	dividerMul := 1
	for _, monkey := range monkeys {
		dividerMul *= monkey.divisibleBy
	}

	return monkeys, dividerMul
}

func parseMonkey(data string) *Monkey {
	parts := strings.Split(data, "\n")

	return &Monkey{
		id:          getId(parts[0]),
		items:       getItems(parts[1]),
		operation:   getOperation(parts[2]),
		divisibleBy: getIntAtPos(parts[3], 21),
		ifTrue:      getIntAtPos(parts[4], 29),
		ifFalse:     getIntAtPos(parts[5], 30),
	}
}

func getId(line string) int {
	id, _ := strconv.Atoi(string(line[len(line)-2]))
	return id
}

func getItems(line string) []int {
	items := []int{}
	ids := strings.Split(line[18:], ", ")
	for _, id := range ids {
		idi, _ := strconv.Atoi(id)
		items = append(items, idi)
	}
	return items
}

func getOperation(line string) func(int) int {
	parts := strings.Split(line[19:], " ")
	if parts[1] == "+" {
		val, _ := strconv.Atoi(parts[2])
		return add(val)
	}
	if parts[2] == "old" {
		return square()
	}
	val, _ := strconv.Atoi(parts[2])
	return mul(val)
}

func add(a int) func(int) int {
	return func(val int) int {
		return val + a
	}
}

func mul(a int) func(int) int {
	return func(val int) int {
		return val * a
	}
}

func square() func(int) int {
	return func(val int) int {
		return val * val
	}
}

func getIntAtPos(line string, pos int) int {
	val, _ := strconv.Atoi(line[pos:])
	return val
}
