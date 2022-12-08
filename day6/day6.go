package day6

// 1909
func PartA(input []byte) any {
	return findMarkerIdx(input, 4)
}

// 3380
func PartB(input []byte) any {
	return findMarkerIdx(input, 14)
}

func findMarkerIdx(data []byte, markerSize int) int {
	lastIdx := map[byte]int{}
	length, start, end := 0, 0, 0
	for {
		start = max(start, lastIdx[data[end]]+1)
		length = max(length, end-start+1)
		lastIdx[data[end]] = end
		if length == markerSize {
			break
		}
		end++
	}
	return end + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
