package ch02

import "fmt"

func isPalindromeUsingStack(root *Node) bool {

	var stack []int

	n := root.len()

	for i := 0; i < n/2; i++ {
		stack = append(stack, root.data)
		root = root.next
	}

	if n%2 == 1 {
		root = root.next
	}

	for root != nil {
		if root.data != stack[len(stack)-1] {
			return false
		}
		stack = stack[0 : len(stack)-1]
		fmt.Println(stack)
		root = root.next
	}

	return true
}

func isPalindromeRecursive(node *Node) {

	list := node.getAsList()
	fmt.Println(isPalindrome(list, 0, len(list)-1))
}

func isPalindrome(arr []int, start, end int) bool {

	//base condition
	if start == end || start > end {
		return true
	}

	if arr[start] != arr[end] {
		return false
	}

	return isPalindrome(arr, start+1, end-1)
}
