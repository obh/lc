package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m < n {
		return findMedian(nums1, nums2)
	} else {
		return findMedian(nums2, nums1)
	}
}

func findMedian(nums1 []int, nums2 []int) float64 {
	fmt.Println(nums1, "and", nums2)
	if len(nums1) <= 2 {
		for i := range nums1 {
			pos := sort.SearchInts(nums2, nums1[i])
			nums2 = insert(nums2, pos, nums1[i])
		}
		_, ans := getMedian(nums2)
		return ans
	}

	position_1, median_1 := getMedian(nums1)
	position_2, median_2 := getMedian(nums2)
	if median_1 == median_2 {
		return median_1
	} else if median_1 < median_2 {
		discarded_len := position_1
		return findMedian(nums1[position_1:], nums2[:len(nums2)-discarded_len])
	} else if median_1 > median_2 {
		evenShift := 0
		if len(nums1)%2 == 0 {
			evenShift += 1
		}
		discarded_len := min(len(nums1)-position_1-1-evenShift, position_2)
		return findMedian(nums1[:len(nums1)-discarded_len], nums2[discarded_len:])
	}
	return 0
}

func getMedian(num []int) (int, float64) {
	m := len(num)
	if m%2 == 0 {
		return m/2 - 1, (float64)(num[m/2]+num[m/2-1]) / 2
	} else {
		return int(m / 2), (float64)(num[m/2])
	}
}

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func main() {
	// fmt.Println(findMedianSortedArrays([]int{1, 2, 4, 8}, []int{3, 5, 7, 9, 11, 15, 19, 20}))
	// fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}))
	fmt.Println(findMedianSortedArrays([]int{4, 5, 6, 8}, []int{1, 2, 3, 7, 9, 10}))
}
