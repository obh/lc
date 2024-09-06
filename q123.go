package main

import (
	"fmt"
)

type point struct {
	value int
	index int
}

func maxProfit(prices []int) int {
	profit := make([]int, len(prices))
	profitIdx := make([]int, len(prices))
	leftQueue := &point{value: prices[0], index: 0}
	for i := 1; i < len(prices); i++ {
		currPrice := prices[i]
		if currPrice <= leftQueue.value {
			profit[i] = 0
			leftQueue.value = currPrice
			leftQueue.index = i
		} else {
			profit[i] = currPrice - leftQueue.value
			profitIdx[i] = leftQueue.index
		}
	}
	fmt.Println("profit if sold at [i]: ", profit)
	// fmt.Println(profitIdx)

	N := len(prices)
	reversePrice := make([]int, len(prices))
	for index, val := range prices {
		reversePrice[N-index-1] = -val
	}
	reverseProfit := make([]int, N)
	revProfitIdx := make([]int, N)
	rightQueue := &point{value: reversePrice[0], index: 0}
	for i := 1; i < N; i++ {
		currPrice := reversePrice[i]
		if currPrice <= rightQueue.value {
			reverseProfit[i] = 0
			rightQueue.value = currPrice
			rightQueue.index = i
		} else {
			reverseProfit[i] = currPrice - rightQueue.value
			revProfitIdx[i] = rightQueue.index
		}
	}
	for left, right := 0, N-1; left < right; left, right = left+1, right-1 {
		reverseProfit[left], reverseProfit[right] = reverseProfit[right], reverseProfit[left]
	}
	fmt.Println("profit if bought at [i]: ", reverseProfit)

	rightMax := make([]int, N)
	for j := N - 2; j >= 1; j-- {
		rightMax[j] = max(rightMax[j+1], reverseProfit[j])
	}
	fmt.Println(rightMax)

	ans := 0
	leftMax := 0
	for i := 0; i <= N-1; i++ {
		leftMax = max(leftMax, profit[i])
		tmp := 0
		if i+1 < N {
			tmp = rightMax[i+1]
		}
		ans = max(ans, leftMax+tmp)
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
