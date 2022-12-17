package day16

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Valve struct {
	name     string
	flowRate int
	tunnels  map[string]int
}

type Valves map[string]Valve

// type Node struct {
// 	name     string
// 	flowRate int
// 	links    []Link
// }

// type Link struct {
// 	from   *Node
// 	to     *Node
// 	weight int
// }

type Player struct {
	current  string
	opened   map[string]bool
	visited  map[string]bool
	minutes  int
	flow     int
	released int
}

func PartA(input []byte) any {
	valves := parseInput(input)
	fmt.Printf("valves: %v\n", len(valves))
	player := newPlayer(valves)
	player.current = "AA"
	player.minutes = 1
	// valves = releaseMaxPressure(valves, player)
	done := false
	for {
		if valves, done = simplify(valves); done {
			break
		}
	}
	fmt.Printf("valves: %v\n", len(valves))

	player = traverse(valves, player)

	return player
}

func PartB(input []byte) any {
	bla := []int{1, 2, 3, 4}
	fmt.Printf("bla: %v\n", bla)
	bla = slices.Delete(bla, 1, 2)
	return bla
}

func traverse(valves Valves, player Player) Player {
	if player.minutes >= 30 {
		fmt.Println("time's up")
		return player
	}
	if allValvesOpen(valves, player) {
		fmt.Println("all valves open")
		return player
	}
	current := valves[player.current]
	player.visited[player.current] = true
	player.released += player.flow
	fmt.Printf("player now here: %v\n", player)

	if canOpen(current, player) {
		fmt.Println("open valve", player.current)
		player.opened[player.current] = true
		player.flow += current.flowRate
		player.minutes += 1
	}
	moveTo := bestTunnel(valves, player)
	fmt.Println("moveTo", moveTo, "takes", current.tunnels[moveTo])
	if moveTo != "" {
		// player.current = moveTo
		// player.minutes += current.tunnels[moveTo]
		player = movePlayerTo(valves, player, moveTo)
		return traverse(valves, player)
	}
	return player
}

// remove inactive valves that only connect two other valves
func simplify(valves Valves) (result Valves, done bool) {
	fmt.Printf("valves: %v\n", valves)
	result = valves
	for _, valve := range valves {
		if valve.name != "AA" && valve.flowRate == 0 && len(valve.tunnels) == 2 {
			valveNames := maps.Keys(valve.tunnels)
			valveNameA, valveNameB := valveNames[0], valveNames[1]
			valveA, valveB := valves[valveNameA], valves[valveNameB]
			delete(valveA.tunnels, valve.name)
			valveA.tunnels[valveB.name] = valve.tunnels[valveNameA] + valve.tunnels[valveNameB]
			delete(valveB.tunnels, valve.name)
			valveB.tunnels[valveA.name] = valve.tunnels[valveNameA] + valve.tunnels[valveNameB]
			delete(valves, valve.name)
			return valves, false
		}
	}
	return valves, true
}

func newPlayer(valves Valves) Player {
	return Player{opened: make(map[string]bool), visited: make(map[string]bool)}
}

func movePlayerTo(valves Valves, player Player, moveTo string) Player {
	time := valves[player.current].tunnels[moveTo]
	player.released += time * player.flow
	player.minutes += time
	player.current = moveTo
	return player
}

func canOpen(valve Valve, player Player) bool {
	return valve.flowRate > 0 && !player.opened[valve.name]
}

func bestTunnel(valves Valves, player Player) string {
	current := valves[player.current]
	maxFlow := 0
	maxTo := ""
	for to, time := range current.tunnels {
		if player.opened[to] && len(valves[to].tunnels) == 1 {
			continue // dead end already visited
		}
		flow := time * player.flow
		if !player.opened[to] {
			flow += valves[to].flowRate
		}
		if flow > maxFlow {
			maxFlow = flow
			maxTo = to
		}
	}
	return maxTo
}

func allValvesOpen(valves Valves, player Player) bool {
	return len(valves) == len(player.opened)
}

func parseInput(input []byte) Valves {
	valves := make(Valves)
	for _, line := range strings.Split(string(input), "\n") {
		regex := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)
		matches := regex.FindStringSubmatch(line)
		tunnels := make(map[string]int)
		for _, link := range strings.Split(matches[3], ", ") {
			tunnels[link] = 1
		}
		valves[matches[1]] = Valve{
			name:     matches[1],
			flowRate: utils.ParseInt(matches[2]),
			tunnels:  tunnels,
		}
	}
	return valves
}
