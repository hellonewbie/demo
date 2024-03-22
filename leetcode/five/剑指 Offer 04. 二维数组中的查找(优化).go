package five

//可以同时考虑一下内部限制和外部限制、、先确定行再确定列
//先考虑本题适用于什么方法
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i = i - 1
		} else if matrix[i][j] < target {
			j = j + 1
		} else {
			return true
			break
		}
	}
	return false
}
