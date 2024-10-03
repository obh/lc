package main

import "fmt"

func candy(ratings []int) int {

	n := len(ratings)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	for i := 0; i < n-1; i++ {
		curr := ratings[i]
		next := ratings[i+1]
		if next > curr && arr[i+1] <= arr[i] {
			arr[i+1] = arr[i] + 1
		}
	}
	for i := n - 1; i > 0; i-- {
		curr := ratings[i]
		prev := ratings[i-1]
		if prev > curr && arr[i-1] <= arr[i] {
			arr[i-1] = arr[i] + 1
		}
	}
	ans := 0
	for i, _ := range arr {
		ans += arr[i]
	}
	return ans
}

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
}
