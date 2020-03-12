package main

import (
	"fmt"
	"math"
)

func main() {
	res := divide(-2147483648, -1)
	fmt.Println("res is ", res)
	fmt.Println("2 32", int(math.Pow(float64(2), float64(31))))
}

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。
示例 1:
输入: dividend = 10, divisor = 3
输出: 3
示例 2:
输入: dividend = 7, divisor = -3
输出: -2
说明:
    被除数和除数均为 32 位有符号整数。
    除数不为 0。
    假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
*/
// 不使用乘法、除法和 mod 运算符,那意思是可以使用加减法，
// 二分法。TODO 这里防止溢出是关键，因为还没有完全掌握位移运算，所以，暂时不管这里
func divide(dividend int, divisor int) int {
	if dividend < 0 && divisor > 0 {
		return -divide(-dividend, divisor)
	}

	if dividend > 0 && divisor < 0 {
		return -divide(dividend, -divisor)
	}

	if dividend < 0 && divisor < 0 {
		return divide(-dividend, -divisor)
	}

	if dividend < divisor {
		return 0
	}
	if dividend == divisor {
		return 1
	}

	result := 1
	p := divisor
	for dividend >= p+p {
		p = p + p
		result = result + result
	}
	result = result + divide(dividend-p, divisor)
	return result
}