package codecn

import "fmt"

//Write a function to find the longest common prefix string amongst an array of
//strings.
//
// If there is no common prefix, return an empty string "".
//
//
// Example 1:
//
//
//Input: strs = ["flower","flow","flight"]
//Output: "fl"
//
//
// Example 2:
//
//
//Input: strs = ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
//
//
//
// Constraints:
//
//
// 0 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lower-case English letters.
//
// Related Topics 字符串
// 👍 1566 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	resStr := ""
	index := 0
	tmp := ""
	for true {
		tmp = strs[0][index : index+1]
		for _, value := range strs {
			if tmp != value[index:index+1] {
				goto endfor
			}
		}
		resStr += tmp
		index++
	}

endfor:
	return resStr
}

func TestP14() {
	strs := []string{"flower", "flow", "flight"}
	resStr := longestCommonPrefix(strs)

	fmt.Println(resStr)
}

//leetcode submit region end(Prohibit modification and deletion)
