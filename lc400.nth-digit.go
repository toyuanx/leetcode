// @Author shovenyuan
// @Date: 2022/6/23 10:48

// Package main ...
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(findNthDigit(11))
}
func findNthDigit(n int) int {
	// 先看数区间
	i := 1
	part := 9 // 区间数字总长度
	num := 0  // n-part
	ans := 0
	for {
		if n <= part {
			num = part - n
			// 计算命中的数
			hitNum := int(math.Pow10(i)) - num/i - 1
			// 计算索引位置
			idx := num % i
			a := strconv.Itoa(hitNum)
			s := a[i-idx-1]
			ans, _ = strconv.Atoi(string(s))
			break
		}
		part += 9 * int(math.Pow10(i)) * (i + 1)
		i++
	}
	// hitNum := num/(i+1) + int(math.Pow10(i))*i
	// a := strconv.Itoa(hitNum)
	// s := a[num%(i+1)]
	// ans, _ := strconv.Atoi(string(s))

	return ans
}
