package main

import "fmt"

func initArr(s string) [52]int32 {
	arr := [52]int32{0}
	for i := 0; i < len(s); i++ {
		ch := int(s[i])
		index := getIndex(ch)
		arr[index] += 1
	}
	return arr
}

func getIndex(ch int) int {
	index := 0
	if ch >= 65 && ch <= 90 {
		index = ch - 65
	} else if ch >= 97 && ch <= 122 {
		index = (ch - 97) + 26
	}
	return index
}

func compare(target [52]int32, src [52]int32) int {
	for i := 0; i < 52; i++ {
		if target[i] > src[i] {
			return -1
		}
	}
	return 0
}

// we know i to j substring has matched. But now we will prune it from left to right
func prune(i int, j int, s string, currArr [52]int32, targetArr [52]int32) (int, int) {
	p := i
	for p = i; p <= j; p++ {
		index := getIndex(int(s[p]))
		if currArr[index] > targetArr[index] {
			currArr[index] -= 1
		} else {
			break
		}
	}
	return p, j
}

// match substring from i
func match(i int, j int, s string, currArr [52]int32, targetArr [52]int32) (int, int, bool) {
	cmp := compare(targetArr, currArr)
	if cmp == 0 {
		tmp1, tmp2 := prune(i, j, s, currArr, targetArr)
		return tmp1, tmp2, true
	}
	//can we add any more characters to the right?
	if j+1 >= len(s) {
		return i, j, false
	}
	newArr := currArr
	index := getIndex(int(s[j+1]))
	newArr[index] += 1
	return match(i, j+1, s, newArr, targetArr)
}

func minWindow(s string, t string) string {

	currArr := initArr(s[0:1])
	targetArr := initArr(t)
	i := 0
	j := 0

	start_index := 0
	substr_length := 1000000000
	for j < len(s) {
		currArr = initArr(s[i : j+1])
		match_i, match_j, isMatched := match(i, j, s, currArr, targetArr)
		if isMatched && (match_j-match_i+1) < substr_length {
			substr_length = (match_j - match_i + 1)
			start_index = match_i
		}
		// fmt.Println("matched -> ", match_i, match_j, isMatched)
		//for next iteration
		i = i + 1
		j = match_j + 1
	}
	if substr_length != 1000000000 {
		return s[start_index : start_index+substr_length]
	} else {
		return ""
	}
}

func main() {
	// target := [52]int32{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// src := [52]int32{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// fmt.Println(compare(target, src))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae")) // -> ab
	// fmt.Println(minWindow("bdab", "ab"))               // -> ab
	// fmt.Println(minWindow("bacdb", "bcd"))             // -> ab
	// fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	// fmt.Println(minWindow("a", "a"))
	// fmt.Println(minWindow("a", "aa"))
}
