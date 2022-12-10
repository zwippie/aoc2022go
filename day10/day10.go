package day10

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

type Cpu struct {
	regX              int
	cycle             int
	signalStrengthSum int
	crt               []bool
}

type Op struct {
	instruction string
	value       int
}

// 17020
func PartA(input []byte) any {
	program := parseInput(input)
	cpu := Cpu{regX: 1, crt: make([]bool, 240)}
	for _, op := range program {
		cpu.execute(op)
	}
	return cpu.signalStrengthSum
}

// RLEZFLGE
func PartB(input []byte) any {
	program := parseInput(input)
	cpu := Cpu{regX: 1, crt: make([]bool, 240)}
	for _, op := range program {
		cpu.execute(op)
	}
	return cpu.crtToString()
}

func (cpu *Cpu) execute(op Op) {
	switch op.instruction {
	case "noop":
		cpu.addCycle()
	case "addx":
		cpu.addCycle()
		cpu.addCycle()
		cpu.regX += op.value
	}
}

func (cpu *Cpu) addCycle() {
	cycle := cpu.cycle % 40
	if cycle == cpu.regX-1 || cycle == cpu.regX || cycle == cpu.regX+1 {
		cpu.crt[cpu.cycle] = true
	} else {
		cpu.crt[cpu.cycle] = false
	}

	cpu.cycle++
	if cpu.cycle%40 == 20 {
		cpu.signalStrengthSum += cpu.cycle * cpu.regX
	}
}

func (cpu *Cpu) crtToString() string {
	result := ""
	for idx, pixel := range cpu.crt {
		if pixel {
			result += "#"
		} else {
			result += "."
		}
		if idx%40 == 39 {
			result += "\n"
		}
	}
	return result
}

func parseInput(input []byte) []Op {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	program := []Op{}

	for scanner.Scan() {
		var line = scanner.Text()
		if line == "noop" {
			program = append(program, Op{instruction: "noop"})
		} else {
			parts := strings.Split(line, " ")
			val, _ := strconv.Atoi(parts[1])
			program = append(program, Op{instruction: "addx", value: val})
		}
	}

	return program
}
