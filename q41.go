package main

import "fmt"

func firstMissingPositive(nums []int) int {
	s := "0"
	count := 100005
	b := make([]byte, len(s)*count)
	bp := 0
	for i := 0; i < count; i++ {
		bp += copy(b[bp:], s)
	}
	for _, n := range nums {
		if n > 0 && n <= count {
			b[n] = 1
		}
	}
	//fmt.Println(b)
	for i := 1; i < count; i++ {
		if b[i] == '0' {
			return i
		}
	}
	return count + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
