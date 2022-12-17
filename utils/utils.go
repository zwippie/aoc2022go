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

func MaxIn(s []int) int {
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func MinIn(s []int) int {
	min := s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func ParseInt(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("%v is not an int: %v", val, err)
	}
	return num
}

func In[T comparable](list []T, val T) bool {
	for _, elem := range list {
		if elem == val {
			return true
		}
	}
	return false
}

func CopyMap[K comparable, V any](original map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range original {
		result[k] = v
	}
	return result
}

// func CopyMap(m map[string]interface{}) map[string]interface{} {
// 	cp := make(map[string]interface{})
// 	for k, v := range m {
// 		vm, ok := v.(map[string]interface{})
// 		if ok {
// 			cp[k] = CopyMap(vm)
// 		} else {
// 			cp[k] = v
// 		}
// 	}
// 	return cp
// }
