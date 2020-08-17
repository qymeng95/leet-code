package blade_offer

func FindRepeatNumber(nums []int) int {
	var t int
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				t = nums[i]
				nums[i] = nums[j]
				nums[j] = t
			}
		}
	}

	for i := 1; i < length; i++ {
		if nums[i-1] == nums[i] {
			return nums[i]
		}
	}

	return -1
}

func FindRepeatNumber2(nums []int) int {
	length := len(nums)
	var numMap = make(map[int]int)
	for i := 0; i < length; i++ {
		if _, ok := numMap[nums[i]]; ok {
			return nums[i]
		} else {
			numMap[nums[i]] = 1
		}
	}

	return -1
}

/**
鸽巢原理，因为出现的元素值 < nums.size(); 所以我们可以将见到的元素 放到索引的位置，如果交换时，发现索引处已存在该元素，则重复 O(N) 空间O(1)
*/
func FindRepeatNumber3(nums []int) int {
	length := len(nums)
	var t int
	for i := 0; i < length; i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}

			t = nums[i]
			nums[i] = nums[t]
			nums[t] = t
		}
	}

	return -1
}

func TestFindRepeatNumber() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	res := FindRepeatNumber(nums)
	res2 := FindRepeatNumber2(nums)
	res3 := FindRepeatNumber3(nums)
	println(res)
	println(res2)
	println(res3)
}
