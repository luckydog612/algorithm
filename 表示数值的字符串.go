package main

import "fmt"

/*
	《剑指offer》面试题19 p127
	 题目：请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
	 例如，字符串“+100”、“5e2”、“-123”、“3.1416”及“-1E-16”都表示数值，
	 但“12e”、“1a3.14”、“1.2.3”、“+-5”及“12e+5.4”都不是。
*/

func isNumber(str string) bool {
	if str == "" {
		return false
	}
	var index = make([]int, 1)
	var isNumber = scanInteger(str, index)

	// 如果出现'.'，则接下来是数字的小数部分
	if index[0] < len(str) && str[index[0]] == '.' {
		index[0]++

		// 下面一行代码用||的原因：
		// 1.小数可以没有整数部分，如.123等于0.123
		// 2.小数点后面可以没有数字，如233.等于233.0
		// 3.当然，小数点前面和后面可以都有数字，如233.666
		isNumber = scanUnsignedInteger(str, index) || isNumber
	}

	// 如果出现'e'或者'E'，则接下来是数字的指数部分
	if index[0] < len(str) && (str[index[0]] == 'e' || str[index[0]] == 'E') {
		index[0]++
		// 下面一行代码用&&的原因：
		// 1.当e或E前面没有数字时，整个字符串不能表示数字，如.e1、e1
		// 2.当e或E后面没有整数时，整个字符串不能表示数字，如12e、12e+5.4
		isNumber = isNumber && scanInteger(str, index)
	}
	return isNumber && index[0] == len(str)
}

func scanInteger(str string, index []int) bool {
	if index[0] < len(str) && (str[index[0]] == '+' || str[index[0]] == '-') {
		index[0]++
	}
	return scanUnsignedInteger(str, index)
}

func scanUnsignedInteger(str string, index []int) bool {
	before := index[0]
	for index[0] < len(str) && str[index[0]] >= '0' && str[index[0]] <= '9' {
		index[0]++
	}
	// 当str中存在若干个0～9的数字时，返回true
	return index[0] > before
}

func main() {
	var num = "12e+5.4"
	fmt.Println(num, isNumber(num))
}
