package offer

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if target == matrix[i][j] {
				return true
			}
		}
	}

	return false
}

func FindNumberIn2DArray2(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	var h = len(matrix)
	var w = len(matrix[0])

	var row = 0
	var col = w - 1

	for row < h && col >= 0 {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return true
		}
	}

	return false
}

func TestFindNumberIn2DArray() {
	nums := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	res := FindNumberIn2DArray(nums, 4)
	res2 := FindNumberIn2DArray2(nums, 4)

	println(res)
	println(res2)
}
