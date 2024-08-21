package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	//sorting the numbers
	h := make(map[string][]int)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum1 := nums[i] + nums[j]
			k := sort.SearchInts(nums, -1*sum1)
			if k < len(nums) && i != j && i != k && j != k && -sum1 == nums[k] {
				// fmt.Println(i, j, k, nums[i], nums[j], nums[k])
				h[getKey(nums[i], nums[j], nums[k])] = []int{nums[i], nums[j], nums[k]}
			}
		}
	}
	ans := make([][]int, 0)
	for _, v := range h {
		ans = append(ans, v)
	}
	return ans
}

func getKey(a, b, c int) string {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	return strconv.Itoa(a) + "_" + strconv.Itoa(b) + "_" + strconv.Itoa(c)
}
