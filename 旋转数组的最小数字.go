package main

import "fmt"

/**
 题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
 例如：数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
*/

func finMinNum(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	mid := start // 如果第一个数字小于最后一个数字，那就说明该数组是有序的，直接返回第一个数字
	for nums[start] >= nums[end] {
		if end-start == 1 {
			mid = end
			break
		}
		mid = (start + end) / 2
		// 如果第一个、最后一个和中间的数字都相等，那么只能按照顺序查找
		if nums[start] == nums[mid] && nums[mid] == nums[end] {
			return finMinInOrder(nums)
		}

		if nums[mid] >= nums[start] {
			start = mid
		} else if nums[mid] <= nums[end] {
			end = mid
		}
	}
	return nums[mid]
}

func finMinInOrder(nums []int) int {
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < num {
			num = nums[i]
		}
	}
	return num
}

func main() {
	nums := []int{1, 1, 1, 1, 0, 1, 1}
	num := finMinNum(nums)
	fmt.Println(num)
}
