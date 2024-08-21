// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 0}
// 	sortColors(arr)
// 	fmt.Println(arr)
// }

// func sortColors(nums []int) {
// 	h := make(map[int]int)
// 	h[0] = 0
// 	h[1] = 0
// 	h[2] = 0
// 	for _, i := range nums {
// 		h[i]++
// 	}
// 	for i := 0; i < h[0]; i++ {
// 		nums[i] = 0
// 	}
// 	for i := h[0]; i < h[0]+h[1]; i++ {
// 		nums[i] = 1
// 	}
// 	for i := h[0] + h[1]; i < h[0]+h[1]+h[2]; i++ {
// 		nums[i] = 2
// 	}
// }
