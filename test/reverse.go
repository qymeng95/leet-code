package test

import "fmt"

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
// 示例 1:
//
// 输入: 123
//输出: 321
//
//
// 示例 2:
//
// 输入: -123
//输出: -321
//
//
// 示例 3:
//
// 输入: 120
//输出: 21
//
//
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// Related Topics 数学

//leetcode submit region begin(Prohibit modification and deletion)
func Reverse(x int) int {
	var res int
	var mod int
	for x != 0 {
		mod = x % 10
		x = x / 10
		res = res*10 + mod
	}

	return res
}

func TestReverse() {
	var res int
	//res = reverse(123)
	//fmt.Print(res)

	res = Reverse(-123)
	fmt.Print(res)
}

//leetcode submit region end(Prohibit modification and deletion)
