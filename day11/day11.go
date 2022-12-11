package day11

import (
	"fmt"
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
	monkeys := parseInput(input)
	for i := 0; i < 20; i++ {
		monkeys.performRound()
		// printMonkeys(monkeys)
	}
	return getScore(monkeys)
}

func PartB(input []byte) any {
	monkeys := parseInput(input)
	for i := 0; i < 1000; i++ {
		monkeys.performRound()
	}
	printMonkeys(monkeys)
	return getScore(monkeys)
}

func (monkeys Monkeys) performRound() {
	for _, monkey := range monkeys {
		// fmt.Printf("monkey in: %v\n", monkey)
		for _, item := range monkey.items {
			item = monkey.operation(item)
			monkey.inspected++
			// item /= 3
			if item%monkey.divisibleBy == 0 {
				otherMonkey := monkeys[monkey.ifTrue]
				// fmt.Printf("otherMonkey in: %v\n", otherMonkey)
				otherMonkey.items = append(otherMonkey.items, item)
				// fmt.Printf("otherMonkey out: %v\n", otherMonkey)
			} else {
				otherMonkey := monkeys[monkey.ifFalse]
				// fmt.Printf("otherMonkey in: %v\n", otherMonkey)
				otherMonkey.items = append(otherMonkey.items, item)
				// fmt.Printf("otherMonkey out: %v\n", otherMonkey)
			}
		}
		monkey.items = []int{}
		// fmt.Printf("monkey out: %v\n", monkey)
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

func printMonkeys(monkeys Monkeys) {
	for _, monkey := range monkeys {
		fmt.Println(monkey)
	}
}

func parseInput(input []byte) Monkeys {
	monkeys := Monkeys{}
	parts := strings.Split(string(input), "\n\n")

	for _, part := range parts {
		monkeys = append(monkeys, parseMonkey(part))
	}

	strings.Split(string(input), "\n\n")

	return monkeys
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
		return mulSelf()
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

func mulSelf() func(int) int {
	return func(val int) int {
		return val * val
	}
}

func getIntAtPos(line string, pos int) int {
	val, _ := strconv.Atoi(line[pos:])
	return val
}
