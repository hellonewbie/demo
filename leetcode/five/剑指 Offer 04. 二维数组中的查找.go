package five

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, list := range row {
			if target == list {
				return true
				break
			}
		}
	}
	return false
}
