package ch10

import (
	"fmt"
	"math/rand"
	"time"
)

func generateArray(n int) []int {

	arr := make([]int, n)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}

	return arr
}

//RankNode represents a node in the binary search tree
type RankNode struct {
	data        int
	left, right *RankNode
	leftsize    int
}

func (rn *RankNode) insert(d int) {

	if d <= rn.data {
		if rn.left == nil {
			rn.left = &RankNode{data: d}
		} else {
			rn.left.insert(d)
		}
		rn.leftsize++

	} else {
		if rn.right == nil {
			rn.right = &RankNode{data: d}
		} else {
			rn.right.insert(d)
		}
	}
}

func (rn *RankNode) getRank(d int) int {

	if rn.data == d {
		return rn.leftsize
	} else if d < rn.data {
		if rn.left == nil {
			return -1
		}
		return rn.left.getRank(d)
	} else {
		if rn.right == nil {
			return -1
		}
		rightrank := rn.right.getRank(d)
		if rightrank == -1 {
			return -1
		}

		return rn.leftsize + 1 + rightrank
	}

}

func inorder(rn *RankNode) {

	if rn != nil {
		inorder(rn.left)
		fmt.Println("data: ", rn.data, ", left: ", rn.leftsize)
		inorder(rn.right)
	}
}
