package day16

import (
	"aoc2022/dijkstra"
	"aoc2022/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/exp/maps"
)

type Valve struct {
	name     string
	flowRate int
	tunnels  map[string]int
}

type Valves map[string]Valve

type Player struct {
	human    string
	elephant string
	opened   map[string]bool
	visited  map[string]bool
	minutes  int
	flow     int
	released int
}

type MaxReleased struct {
	value int
	mutex sync.RWMutex
}

var maxReleased *MaxReleased

func updateMaxReleased(mr *MaxReleased, value int) {
	mr.mutex.RLock()
	defer mr.mutex.RUnlock()
	mr.value = utils.Max(mr.value, value)
}

// 1751
func PartA(input []byte) any {
	valves := parseInput(input)
	fmt.Printf("valves: %v\n", len(valves))
	valves = simplifyValves(valves)
	fmt.Printf("valves: %v\n", len(valves))

	player := newPlayer(valves)
	maxReleased = &MaxReleased{value: math.MinInt}

	graph := buildGraph(valves)
	fmt.Printf("graph: \n%v\n", graph)

	traverse(valves, player, precalculateDistances(valves))

	return maxReleased.value
}

func PartB(input []byte) any {
	valves := parseInput(input)
	fmt.Printf("valves: %v\n", len(valves))
	valves = simplifyValves(valves)
	fmt.Printf("valves: %v\n", len(valves))

	// player := newPlayer(valves)
	maxReleased = &MaxReleased{value: math.MinInt}

	graph := buildGraph(valves)
	fmt.Printf("graph: \n%v\n", graph)

	// traverseWithElephant(valves, player, precalculateDistances(valves))

	return maxReleased.value
}

// func traverseWithElephant(valves Valves, player *Player, distances map[string]map[string]int) {
// 	humanValve := valves[player.human]
// 	elephantValve := valves[player.elephant]

// 	player.visited[player.human] = true
// 	player.visited[player.elephant] = true
// 	player.released += player.flow

// 	if player.minutes >= 24 {
// 		// time's up
// 		if player.minutes == 24 {
// 			updateMaxReleased(maxReleased, player.released)
// 		}
// 		return
// 	}

// 	player.minutes += 1

// 	if allValvesOpen(valves, player) {
// 		// just wait
// 		traverseWithElephant(valves, player, distances)
// 		return
// 	}

// 	if canOpen(current, player) {
// 		// open valve
// 		player.opened[player.human] = true
// 		player.flow += current.flowRate
// 		traverse(valves, player, distances)
// 	} else {
// 		// try any next non-opened valve
// 		for _, valve := range valves {
// 			if valve.name != player.human && canOpen(valve, player) {
// 				// fly to valve
// 				minutes := distances[player.human][valve.name] - 1 // already added 1
// 				playerCopy := copyPlayer(player)
// 				playerCopy.human = valve.name
// 				playerCopy.minutes += minutes
// 				playerCopy.released += (minutes * playerCopy.flow)
// 				traverse(valves, playerCopy, distances)
// 			}
// 		}
// 	}
// }

// Depth First Search while keeping track of global max released. Not fast.
func traverse(valves Valves, player *Player, distances map[string]map[string]int) {
	current := valves[player.human]
	player.visited[player.human] = true
	player.released += player.flow

	if player.minutes >= 30 {
		// time's up
		if player.minutes == 30 {
			updateMaxReleased(maxReleased, player.released)
		}
		return
	}

	player.minutes += 1

	if allValvesOpen(valves, player) {
		// just wait
		traverse(valves, player, distances)
		return
	}

	if canOpen(current, player) {
		// open valve
		player.opened[player.human] = true
		player.flow += current.flowRate
		traverse(valves, player, distances)
	} else {
		// try any next non-opened valve
		for _, valve := range valves {
			if valve.name != player.human && canOpen(valve, player) {
				// fly to valve
				minutes := distances[player.human][valve.name] - 1 // already added 1
				playerCopy := copyPlayer(player)
				playerCopy.human = valve.name
				playerCopy.minutes += minutes
				playerCopy.released += (minutes * playerCopy.flow)
				traverse(valves, playerCopy, distances)
			}
		}
	}
}

func precalculateDistances(valves Valves) map[string]map[string]int {
	result := make(map[string]map[string]int)
	for _, valve := range valves {
		graph := buildGraph(valves)
		dijkstra.Dijkstra(graph, valve.name)
		result[valve.name] = make(map[string]int)
		for _, valveB := range valves {
			result[valve.name][valveB.name] = graph.GetNode(valveB.name).Value
		}
	}
	return result
}

// repeatedly remove inactive valves
func simplifyValves(valves Valves) Valves {
	done := false
	for {
		if valves, done = simplify(valves); done {
			break
		}
	}
	return valves
}

// remove inactive valves that only connect two other valves
func simplify(valves Valves) (result Valves, done bool) {
	// fmt.Printf("valves: %v\n", valves)
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

func newPlayer(valves Valves) *Player {
	player := &Player{
		human:   "AA",
		minutes: 1,
		opened:  make(map[string]bool),
		visited: make(map[string]bool),
	}
	player.opened["AA"] = true // hack to make allValvesOpen simpler as AA is the only 0-valve
	return player
}

func copyPlayer(player *Player) *Player {
	return &Player{
		human:    player.human,
		opened:   utils.CopyMap(player.opened),
		visited:  utils.CopyMap(player.visited),
		minutes:  player.minutes,
		flow:     player.flow,
		released: player.released,
	}
}

func canOpen(valve Valve, player *Player) bool {
	return valve.flowRate > 0 && !player.opened[valve.name]
}

func allValvesOpen(valves Valves, player *Player) bool {
	return len(valves) == len(player.opened)
}

func buildGraph(valves Valves) *dijkstra.WeightedGraph {
	graph := dijkstra.NewGraph()
	for _, valve := range valves {
		graph.AddNode(&dijkstra.Node{Name: valve.name, Value: math.MaxInt, Through: nil})
	}
	for _, valve := range valves {
		from := graph.GetNode(valve.name)
		for toName, distance := range valve.tunnels {
			to := graph.GetNode(toName)
			graph.AddEdge(from, to, distance)
		}
	}
	return graph
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
