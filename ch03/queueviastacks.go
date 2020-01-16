package ch03

//stack implementation starts here
type customstack struct {
	items []int
}

func newCustomStack() customstack {

	s := customstack{}
	s.items = make([]int, 0)
	return s
}

func (s customstack) push(i int) customstack {

	s.items = append(s.items, i)
	return s
}

func (s customstack) pop() (int, bool, customstack) {

	if s.isEmpty() {
		return 0, false, s
	}

	deleted := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return deleted, true, s

}

func (s customstack) peek() (int, bool) {

	if s.isEmpty() {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

func (s customstack) isEmpty() bool {
	return len(s.items) == 0
}

//stack implementation ends here

//queue implementation starts here
type customqueue struct {
	stack1 customstack
	stack2 customstack
}

func newCustomQueue() customqueue {

	q := customqueue{}
	q.stack1 = newCustomStack()
	q.stack2 = newCustomStack()

	return q
}

func (q customqueue) add(i int) customqueue {

	q.stack1.items = append(q.stack1.items, i)
	return q
}

func (q customqueue) remove() customqueue {

	if q.stack1.isEmpty() {
		return q
	}

	var deleted int

	for !q.stack1.isEmpty() {

		deleted, _, q.stack1 = q.stack1.pop()
		q.stack2 = q.stack2.push(deleted)
	}

	_, _, q.stack2 = q.stack2.pop()

	for !q.stack2.isEmpty() {

		deleted, _, q.stack2 = q.stack2.pop()
		q.stack1 = q.stack1.push(deleted)
	}

	return q

}
