package day8

// 1700
func PartA(input []byte) any {
	data := parseInput(input)
	rowLen, colLen := len(data), len(data[0])
	visibleTreeCount := 2*(rowLen-1) + 2*(colLen-1)

	for rowIdx := 1; rowIdx < rowLen-1; rowIdx++ {
		for colIdx := 1; colIdx < colLen-1; colIdx++ {
			value := data[rowIdx][colIdx]
			// left smaller? right smaller? up smaller? down smaller?
			leftVisible := true
			for colIdx2 := 0; colIdx2 < colIdx; colIdx2++ {
				if data[rowIdx][colIdx2] >= value {
					leftVisible = false
					break
				}
			}
			rightVisible := true
			for colIdx2 := colIdx + 1; colIdx2 < colLen; colIdx2++ {
				if data[rowIdx][colIdx2] >= value {
					rightVisible = false
					break
				}
			}
			upVisible := true
			for rowIdx2 := 0; rowIdx2 < rowIdx; rowIdx2++ {
				if data[rowIdx2][colIdx] >= value {
					upVisible = false
					break
				}
			}
			downVisible := true
			for rowIdx2 := rowIdx + 1; rowIdx2 < rowLen; rowIdx2++ {
				if data[rowIdx2][colIdx] >= value {
					downVisible = false
					break
				}
			}
			if leftVisible || rightVisible || upVisible || downVisible {
				visibleTreeCount++
			}
		}
	}

	return visibleTreeCount
}

// 470596
func PartB(input []byte) any {
	data := parseInput(input)
	rowLen, colLen := len(data), len(data[0])
	maxScenicScore := 0

	for rowIdx := 0; rowIdx < rowLen; rowIdx++ {
		for colIdx := 0; colIdx < colLen; colIdx++ {
			maxScenicScore = max(maxScenicScore, scenicScore(data, rowIdx, colIdx))
		}
	}

	return maxScenicScore
}

func scenicScore(data [][]int, rowIdx int, colIdx int) int {
	rowLen, colLen := len(data), len(data[0])
	value := data[rowIdx][colIdx]

	// look up
	upScore := 0
	for rowIdx2 := rowIdx - 1; rowIdx2 >= 0; rowIdx2-- {
		upScore++
		if data[rowIdx2][colIdx] >= value {
			break
		}
	}
	// loop down
	downScore := 0
	for rowIdx2 := rowIdx + 1; rowIdx2 < rowLen; rowIdx2++ {
		downScore++
		if data[rowIdx2][colIdx] >= value {
			break
		}
	}
	// look left
	leftScore := 0
	for colIdx2 := colIdx - 1; colIdx2 >= 0; colIdx2-- {
		leftScore++
		if data[rowIdx][colIdx2] >= value {
			break
		}
	}
	// look right
	rightScore := 0
	for colIdx2 := colIdx + 1; colIdx2 < colLen; colIdx2++ {
		rightScore++
		if data[rowIdx][colIdx2] >= value {
			break
		}
	}

	return upScore * downScore * leftScore * rightScore
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInput(input []byte) [][]int {
	result := [][]int{}
	row := []int{}
	for _, char := range input {
		if char == 10 {
			result = append(result, row)
			row = []int{}
		} else {
			row = append(row, int(char)-48)
		}
	}
	return append(result, row)
}
