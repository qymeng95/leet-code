package main

import "fmt"

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//
//

// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	var result []int

	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)

				return result
			}
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func testTwoSum() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Print(result)
}
