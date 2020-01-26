package ch04

func getNoOfPathsFromNode(root *Node, target int, currentsum int) int {

	if root == nil {

		return 0
	}

	currentsum += root.data

	totalPaths := 0

	if currentsum == target {
		totalPaths++
	}

	totalPaths += getNoOfPathsFromNode(root.left, target, currentsum)
	totalPaths += getNoOfPathsFromNode(root.right, target, currentsum)

	return totalPaths
}

func getNoOfPaths(root *Node, target int) int {

	if root == nil {
		return 0
	}

	fromRoot := getNoOfPathsFromNode(root, target, 0)
	fromLeft := getNoOfPaths(root.left, target)
	fromRight := getNoOfPaths(root.right, target)

	return fromRoot + fromLeft + fromRight

}
