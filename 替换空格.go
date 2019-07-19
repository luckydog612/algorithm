package main

import "fmt"

/*
	请实现一个函数，把字符串中的每个空格替换成"%20"。
	例如，输入“We are happy.”，则输出“We%20are%20happy.”。
*/

// 方法一：通常我们会想到遍历数组，遇到空格时，进行替换，移动后面的字符，这样会比较麻烦，而且效率低
// 方法二：反向思维，从后面往前遍历，移动字符，遇到空格替换
func ReplaceSpace(str string, fillStr string) string {
	subLen := len(fillStr)
	strLen := len(str)
	if subLen == 0 || strLen == 0 {
		return ""
	}

	//1.统计空格的个数，计算替换后字符串的总长度
	count := 0
	for i := 0; i < strLen; i++ {
		if str[i] == ' ' {
			count++
		}
	}
	if count == 0 {
		return str
	}
	//2.初始化一个足够长的切片
	length := count*subLen - count
	s := make([]byte, length)
	ss := append(StringTiArray(str), s...)
	//3.遍历数组，移动非空格字符
	var index = strLen - 1   // 标记字符串的末尾
	var index2 = len(ss) - 1 // 标记末尾处的空位
	for i := strLen - 1; i >= 0; i-- {
		if ss[i] == ' ' {
			fillSbuStr(ss, index2, fillStr)
			index2 -= subLen
			index--
		} else {
			ss[index2] = ss[index]
			index--
			index2--
		}
	}
	return string(ss)
}

// 将字符串转化为字节数组
func StringTiArray(str string) []byte {
	arr := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		arr[i] = str[i]
	}
	return arr
}

func fillSbuStr(s []byte, index int, subStr string) {
	for j := len(subStr) - 1; j >= 0; j-- {
		s[index] = subStr[j]
		index--
	}
}

func main() {
	var str = " We are happy. "
	var subStr = "%20"
	newStr := ReplaceSpace(str, subStr)
	fmt.Println(str)
	fmt.Println(newStr)
}
