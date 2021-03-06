package main

import (
	"fmt"
)

func main() {
	gas := []int{3, 3, 4}
	cost := []int{3, 4, 4}
	// gas := []int{5, 1, 2, 3, 4}
	// cost := []int{4, 4, 1, 5, 1}
	res := canCompleteCircuit(gas, cost)
	fmt.Println("res is ", res)
}

/*
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
说明:
    如果题目有解，该答案即为唯一答案。
    输入数组均为非空数组，且长度相同。
    输入数组中的元素均为非负数。
示例 1:
输入:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
输出: 3
链接：https://leetcode-cn.com/problems/gas-station
*/
func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(cost) == 0 || len(gas) != len(cost) {
		return -1
	}
	currGas, totalGas, start := 0, 0, 0
	n := len(gas)
	for i := 0; i < n; i++ {
		totalGas += gas[i] - cost[i]
		currGas += gas[i] - cost[i]
		if currGas < 0 {
			start = i + 1
			currGas = 0
		}
	}
	if totalGas >= 0 {
		return start
	}
	return -1
}
