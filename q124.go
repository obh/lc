package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SumNode struct {
	LeftSum      int
	RightSum     int
	LeftRightSum int
	Ans          int
}

func maxPathSum(root *TreeNode) int {
	sumNode := computeLargestSum(root)
	return sumNode.Ans
}

func computeLargestSum(node *TreeNode) *SumNode {
	if node == nil {
		return &SumNode{
			LeftSum:      -1000000,
			RightSum:     -1000000,
			LeftRightSum: -1000000,
			Ans:          -1000000,
		}
	}

	leftSumNode := computeLargestSum(node.Left)
	rightSumNode := computeLargestSum(node.Right)

	leftSum := getMax(leftSumNode.LeftSum+node.Val, leftSumNode.RightSum+node.Val, node.Val)
	rightSum := getMax(rightSumNode.LeftSum+node.Val, rightSumNode.RightSum+node.Val, node.Val)
	leftRightSum :=
		Max(leftSumNode.LeftSum, leftSumNode.RightSum) + Max(rightSumNode.LeftSum, rightSumNode.RightSum) + node.Val

	Ans := getMax(leftSum, rightSum, leftRightSum, leftSumNode.Ans, rightSumNode.Ans)

	tmp := &SumNode{
		LeftSum:      leftSum,
		RightSum:     rightSum,
		LeftRightSum: leftRightSum,
		Ans:          Ans,
	}
	fmt.Printf("returning for node: %d, %+v\n", node.Val, tmp)
	return tmp
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMax(arr ...int) int {
	max := -1000000
	for _, a := range arr {
		if a > max {
			max = a
		}
	}
	return max
}

func main() {
	// left := TreeNode{Val: 3, Left: nil, Right: nil}
	// right := TreeNode{Val: 3, Left: nil, Right: nil}
	// root := TreeNode{Val: 1, Left: &left, Right: &right}
	// fmt.Println(maxPathSum(&root))

	// t1_l1 := TreeNode{Val: 9, Left: nil, Right: nil}
	// t1_l2 := TreeNode{Val: 15, Left: nil, Right: nil}
	// t1_l3 := TreeNode{Val: 7, Left: nil, Right: nil}
	// t1_l4 := TreeNode{Val: 20, Left: &t1_l2, Right: &t1_l3}
	// t1_l5 := TreeNode{Val: -10, Left: &t1_l1, Right: &t1_l4}
	// fmt.Println(maxPathSum(&t1_l5))

	t2_l1 := TreeNode{Val: 4, Left: nil, Right: nil}
	t2_l2 := TreeNode{Val: -1, Left: nil, Right: &t2_l1}
	t2_l3 := TreeNode{Val: 1, Left: nil, Right: nil}
	t2_l4 := TreeNode{Val: -2, Left: &t2_l3, Right: &t2_l2}
	t2_l5 := TreeNode{Val: 5, Left: nil, Right: &t2_l4}
	fmt.Println(maxPathSum(&t2_l5))
}
