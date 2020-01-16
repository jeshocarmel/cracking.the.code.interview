package ch03

type node struct {
	data int
	min  int
}

type minstack struct {
	items []node
}

func newMinStack() minstack {

	s := minstack{}
	s.items = make([]node, 0)
	return s
}

func (ms minstack) push(i int) minstack {

	newnode := node{data: i}
	if ms.isEmpty() {
		newnode.min = i
	} else {

		top := ms.peek()
		if i < top.min {
			newnode.min = i
		} else {
			newnode.min = top.min
		}
	}

	ms.items = append(ms.items, newnode)
	return ms
}

func (ms minstack) pop() minstack {

	if ms.isEmpty() {
		return ms
	}

	ms.items = ms.items[0 : len(ms.items)-1]
	return ms
}

func (ms minstack) peek() *node {

	if ms.isEmpty() {
		return nil
	}

	return &ms.items[len(ms.items)-1]
}

func (ms minstack) isEmpty() bool {
	return len(ms.items) == 0
}

func (ms minstack) min() int {
	return ms.peek().min
}
