package main

import "fmt"

func numDistinct(s string, t string) int {
	m := len(s)

	a := make([][]int, m)

	for i := range s {
		a[i] = make([]int, len(t))
	}

	/*
		a[i][j] = a[k][j] where k is max[0, i-1] st a[k][j] > 0 +
				  a[k][j-1] where ks is max[0, i-1] st a[k][j-1] > 0

	*/
	currCol := 0
	prevCol := 0
	// zerothCol := make([]int, len(s))
	for j, destChar := range t {
		for i, srcChar := range s {
			if j > i { //not possible [ignore]
				continue
			}
			//update
			if j == 0 {
				prevCol = 1
				// zerothCol[i-1] = a[i-1][j]
			} else {
				prevCol = max(prevCol, a[i-1][j-1])
			}
			// if i == 2 {
			// 	fmt.Printf("a[%d][%d]: %s == %s?  (%d + %d) AND A[i-1][j]=%d\n", i, j, string(srcChar), string(destChar), currCol, prevCol, a[i-1][j])
			// }
			if srcChar == destChar {
				a[i][j] += currCol
				a[i][j] += prevCol
				currCol = max(a[i][j], currCol)
			}

		}
		currCol = 0
		prevCol = 0
	}
	fmt.Println(a)
	ans := 0
	for i := range s {
		ans = max(ans, a[i][len(t)-1])
	}
	return ans
}

func main() {
	fmt.Println(numDistinct("rat", "ra"))
	fmt.Println(numDistinct("babgbag", "bag"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("p", "qq"))

}
