package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	zeros := strings.Repeat("0", 60010)
	for _, i := range nums {
		index := 30000 + i
		if zeros[index] == '0' {
			zeros = zeros[:index] + "1" + zeros[index+1:]
		} else {
			zeros = zeros[:index] + "0" + zeros[index+1:]
		}
	}
	return strings.Index(zeros, "1") - 30000
}
