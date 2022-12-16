package utils

import (
	"log"
	"strconv"
)

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func ParseInt(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("%v is not an int: %v", val, err)
	}
	return num
}
