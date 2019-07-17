package main

import "fmt"

/*
	在一个长度为n的数组里的所有数字都在0～n-1的范围内。
	数组中某些数字是重复的，但不知道有几个数字重复了，
	也不知道了每个数字重复了几次。请找出数组中任意一个重复的数字。
	例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是重复的数字2或者3。
*/

// 方法一：最简单的是直接使用哈希表来找到，这里我们用map
func findTheRepeatNum(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exist := m[nums[i]]; exist {
			return nums[i]
		}
		m[nums[i]] = 1
	}
	return -1
}

// 方法二：利用数组的下标
// 遍历数组，当扫描到下标为i的数字时，首先比较这个数字m是不是等于i。
// 如果相等，则接着扫描下一个数字；如果不是则再拿这个数字m和第m个数字是否相等，
// 如果相等，就找到了一个重复的数字；如果不相等，第i个和第m个位置的数字交换，继续比较。
func findTheRepeatNum2(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}

	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}

	for i := 0; i < len(nums); i++ {
		m := nums[i]
		if m != i {
			if nums[m] == m {
				return m
			} else {
				nums[i], nums[m] = nums[m], nums[i]
			}
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num1 := findTheRepeatNum(nums)
	num2 := findTheRepeatNum2(nums)
	fmt.Println(num1, num2)
}
