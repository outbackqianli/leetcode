package main

import (
	"fmt"
	"math"
)

func main() {
	res := poorPigs(1, 1, 1)
	fmt.Println("res is ", res)
}

/*
有 1000 只水桶，其中有且只有一桶装的含有毒药，其余装的都是水。它们从外观看起来都一样。如果小猪喝了毒药，它会在 15 分钟内死去。
问题来了，如果需要你在一小时内，弄清楚哪只水桶含有毒药，你最少需要多少只猪？
回答这个问题，并为下列的进阶问题编写一个通用算法。
进阶:
假设有 n 只水桶，猪饮水中毒后会在 m 分钟内死亡，你需要多少猪（x）就能在 p 分钟内找出 “有毒” 水桶？
这 n 只水桶里有且仅有一只有毒的桶。
提示：
    可以允许小猪同时饮用任意数量的桶中的水，并且该过程不需要时间。
    小猪喝完水后，必须有 m 分钟的冷却时间。在这段时间里，只允许观察，而不允许继续饮水。
    任何给定的桶都可以无限次采样（无限数量的猪）。
*/

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets == 0 || minutesToDie == 0 || minutesToTest == 0 || minutesToTest < minutesToDie {
		return 0
	}
	bit := minutesToTest/minutesToDie + 1

	m := 0
	pig := 0 // 也不知道测试用例是怎么弄的,1,1,1的输入时,答案是0而不是1,所以这里就只能是0

	for {
		m = int(math.Pow(float64(bit), float64(pig)))
		//fmt.Println("m pig ", bit, m, pig)
		if m >= buckets {
			return pig
		}
		pig++
	}
}
