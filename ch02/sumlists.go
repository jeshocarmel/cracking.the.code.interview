package ch02

func sumListsReverse(list1, list2 *Node) int {

	count := 1
	num1 := 0
	for n := list1; n != nil; n = n.next {
		num1 += count * (n.data % 10)
		count = count * 10
	}

	count = 1
	num2 := 0

	for n := list2; n != nil; n = n.next {
		num2 += count * (n.data % 10)
		count = count * 10
	}

	return num1 + num2
}

func sumListsNormal(list1, list2 *Node) int {

	len1 := list1.len()
	len2 := list2.len()

	factor := 1

	for i := 1; i < len1; i++ {
		factor *= 10
	}

	num1 := 0
	for n := list1; n != nil; n = n.next {
		num1 += n.data * factor
		factor = factor / 10
	}

	factor = 1
	for i := 1; i < len2; i++ {
		factor *= 10
	}

	num2 := 0
	for n := list2; n != nil; n = n.next {
		num2 += n.data * factor
		factor = factor / 10
	}

	return num1 + num2

}
