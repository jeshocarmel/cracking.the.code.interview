package ch04

import (
	"strconv"
	"strings"
)

//Node represents a tree Node
type Node struct {
	data  int
	left  *Node
	right *Node
}

func createNode(data int) *Node {
	return &Node{data: data}
}

func insert(root *Node, data int) *Node {
	if root == nil {
		return createNode(data)
	}

	if data < root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}

	return root
}

var sb strings.Builder

func inOrder(root *Node) string {

	if root == nil {
		return ""
	}

	inOrder(root.left)
	sb.WriteString(strconv.Itoa(root.data))
	sb.WriteRune(' ')
	inOrder(root.right)

	return sb.String()
}

func divideAndConquer(arr []int, start, end int, root *Node) *Node {

	if start >= end {
		return nil
	}

	mid := (start + end) / 2

	root = insert(root, arr[mid])

	divideAndConquer(arr, start, mid, root)
	divideAndConquer(arr, mid+1, end, root)

	return root
}
