package main

import "fmt"

/*
	《剑指offer》面试题17 p114
	 题目：输入数字n，按顺序打印从1到最大的n位十进制数。
	比如输入3，则打印出1,2,3一直到最大的3位数999
*/

func Print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}

	number := make([]int, n)
	for i := 0; i < 10; i++ {
		number[0] = i
		print1ToMaxOfDigitsRecursively(number, n, 0)
	}
}

func print1ToMaxOfDigitsRecursively(number []int, length int, index int) {
	if index == length-1 {
		printNumber(number)
		return
	}

	for i := 0; i < 10; i++ {
		number[index+1] = i
		print1ToMaxOfDigitsRecursively(number, length, index+1)
	}
}

func printNumber(number []int) {
	var isBeginning0 = true
	length := len(number)
	for i := 0; i < length; i++ {
		if isBeginning0 && number[i] != 0 {
			isBeginning0 = false
		}

		if !isBeginning0 {
			fmt.Printf("%d", number[i])
			if i == length-1 {
				fmt.Println()
			}
		}
	}
}

func main() {
	Print1ToMaxOfDigits(5)
}
