package day25

import (
	"aoc2022/utils"
	"strings"
)

// 2-2--02=1---1200=0-1
func PartA(input []byte) any {
	sum := 0
	for _, val := range parseInput(input) {
		decimal := snafuToDecimal(val)
		sum += decimal
		// fmt.Printf("snafuToDecimal(%v): %v\n", val, decimal)
		// snafu := decimalToSnafu(decimal)
		// fmt.Printf("snafu: %v\n", snafu)
		// if val != snafu {
		// 	log.Fatal("err")
		// }
	}
	return decimalToSnafu(sum)
}

func PartB(input []byte) any {
	return 0
}

func snafuToDecimal(val string) int {
	runes := []rune(val)
	result := 0
	for i, r := range runes {
		p := len(runes) - i - 1
		if r == '=' {
			result -= (2 * pow5(p))
		} else if r == '-' {
			result -= pow5(p)
		} else if r == '1' {
			result += pow5(p)
		} else if r == '2' {
			result += (2 * pow5(p))
		}
	}
	return result
}

func decimalToSnafu(val int) string {
	digits := 0
	for {
		if val <= maxWithDigits(digits) {
			break
		}
		digits++
	}

	result := ""
	for d := digits; d > 0; d-- {
		max := maxWithDigits(d - 1)
		r := "0"
		if val > 0 {
			if val-pow5(d-1) >= -max {
				val -= pow5(d - 1)
				r = "1"
			}
			if val-pow5(d-1) >= -max {
				val -= pow5(d - 1)
				r = "2"
			}
		} else if val < 0 {
			if val+pow5(d-1) <= max {
				val += pow5(d - 1)
				r = "-"
			}
			if val+pow5(d-1) <= max {
				val += pow5(d - 1)
				r = "="
			}
		}
		result += r
	}

	return result
}

func maxWithDigits(digits int) int {
	result := 0
	for d := 0; d < digits; d++ {
		result += 2 * pow5(d)
	}
	return result
}

func pow5(val int) int {
	return utils.Pow(5, val)
}

func parseInput(input []byte) []string {
	return strings.Split(string(input), "\n")
}
