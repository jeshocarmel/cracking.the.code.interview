package ch04

import (
	"math"
)

func (n *Node) height() int {

	if n == nil {
		return -1
	}

	leftHeight := n.left.height()
	if leftHeight == math.MinInt8 {
		return math.MinInt8
	}

	rightHeight := n.right.height()
	if rightHeight == math.MinInt8 {
		return math.MinInt8
	}

	heightDiff := abs(leftHeight - rightHeight)

	if heightDiff > 1 {
		return math.MinInt8
	}
	return max(leftHeight, rightHeight) + 1
}

func (n *Node) isBalanced() bool {

	return n.height() != math.MinInt8
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
