package ch04

import "math"

func isBalancedBruteForce(root *Node) bool {

	if root == nil {
		return true
	}

	lHeight := getHeight(root.left)
	rHeight := getHeight(root.right)

	if abs(lHeight-rHeight) > 1 {
		return false
	}

	return isBalancedBruteForce(root.left) && isBalancedBruteForce(root.right)

}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func getHeight(root *Node) int {

	if root == nil {
		return 0
	}

	lheight := getHeight(root.left)
	rheight := getHeight(root.right)

	if lheight > rheight {
		return lheight + 1
	}

	return rheight + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func checkHeights(root *Node) (int, bool) {

	if root == nil {
		return 0, true
	}

	lheight, isLeftValid := checkHeights(root.left)
	if !isLeftValid {
		return math.MaxInt8, false
	}
	rheight, isRightValid := checkHeights(root.right)
	if !isRightValid {
		return math.MaxInt8, false
	}

	if abs(lheight-rheight) > 1 {
		return math.MaxInt8, false
	}

	if lheight > rheight {
		return lheight + 1, true
	}

	return rheight + 1, true

}

func isBalancedOptimized(root *Node) bool {

	if root == nil {
		return true
	}

	_, res := checkHeights(root)

	return res

}
