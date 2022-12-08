package main

import "fmt"

type Node struct {
	Height int
	Width  int
}

func trap(height []int) int {
	ans := 0
	nodes := []Node{}
	for _, h := range height {
		nodes = append(nodes, Node{h, 1})
	}
	//fmt.Println(maxToLeft)
	//0 and 1 can be ingored
	for i := 2; i < len(nodes); {
		anchor := i
		peak := findPeak(nodes, i)
		//has to have atleast something between these two
		//since nodes is dynamic array we have to manage increment by ourselves
		if anchor-peak >= 2 {
			// debug(nodes, peak, anchor, ans, 0)
			minHt := min(nodes[peak].Height, nodes[anchor].Height)
			totalWidth := 0
			thisStorage := 0
			for j := peak + 1; j <= anchor-1; j++ {
				//only if there is space can you store water
				if minHt > nodes[j].Height {
					thisStorage += (minHt - nodes[j].Height) * nodes[j].Width
				}
				totalWidth += nodes[j].Width
			}
			if thisStorage > 0 {
				nodes[peak+1].Height = minHt
				nodes[peak+1].Width = totalWidth
				ans += thisStorage
			}
			//squish if anchor - peak >= 3
			if anchor-peak >= 3 && thisStorage > 0 {
				//fmt.Println("HELLO WORLD, SQUISH PLEASE")
				nodes = append(nodes[:peak+2], nodes[anchor:]...)
				i = peak + 2
			}
			// debug(nodes, peak, i, ans, 1)
		}
		//always increment
		i += 1
	}
	return ans
}

func findPeak(nodes []Node, end int) int {
	currHeight := -1
	currIndex := 0
	for i := end - 1; i >= 0; i-- {
		if nodes[i].Height >= nodes[end].Height {
			return i
		}
		if nodes[i].Height > currHeight {
			currIndex = i
			currHeight = nodes[i].Height
		}
	}
	return currIndex
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap([]int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 1, 2, 1, 3, 4, 6, 8, 1, 3}))
}
