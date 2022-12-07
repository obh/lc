package main

import "fmt"

// ( -> 0
// ) -> -1
func longestValidParentheses(s string) int {
	q := []int{}
	for i := 0; i < len(s); i++ {
		//setting the char
		char := 0
		if s[i] == ')' {
			char = -1
		}

		//this depicts the length of the longest substring
		length := 0
		for j := len(q) - 1; j >= 0; j-- {
			if q[j] > 0 {
				length += q[j]
				continue
			}
			//match
			if q[j] == 0 && char == -1 {
				q[j] = length + 2
				q = q[:j+1]
				char = 1 //denotes that char is handled
			} else if q[j] == -1 && char == -1 {
				q = append(q, char)
				char = 1
			} else if char == 0 {
				q = append(q, char)
				char = 1
			}
		}
		//not handled
		if char != 1 {
			q = append(q, char)
		}
	}

	ans := 0
	runningCount := 0
	//fmt.Println(q)
	for i := 0; i < len(q); i++ {
		if q[i] > 1 {
			runningCount += q[i]
		}
		ans = max(runningCount, ans)
		if q[i] <= 1 {
			runningCount = 0
		}
	}
	if ans == 1 {
		return 0
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestValidParentheses("()())"))
	fmt.Println(longestValidParentheses("((()(())"))
	fmt.Println(longestValidParentheses("((()(())(())))"))
	fmt.Println(longestValidParentheses("((()(())(())))()("))
	fmt.Println(longestValidParentheses("(()"))
}
