package ch03

type stack struct {
	capacity       int
	items          []int
	numberofstacks int
	sizes          []int
}

func newStack(nos, totalcapacity int) stack {
	s := stack{numberofstacks: nos, capacity: totalcapacity}
	s.items = make([]int, totalcapacity)
	s.sizes = make([]int, nos)
	return s
}

func (s stack) push(stackno, data int) stack {

	if s.sizes[stackno] >= s.getIndividualCap() {
		return s
	}

	index := s.getIndexOfTop(stackno)
	s.items[index+1] = data
	s.sizes[stackno]++

	return s
}

func (s stack) pop(stackno int) (stack, int, bool) {

	if s.isEmpty(stackno) {
		return s, 0, false
	}

	topIndex := s.getIndexOfTop(stackno)
	deleted := s.items[topIndex]

	s.items[topIndex] = 0
	s.sizes[stackno]--

	return s, deleted, true

}

func (s stack) isEmpty(stackno int) bool {
	return s.sizes[stackno] == 0
}

func (s stack) getIndexOfTop(stackno int) int {

	offset := s.getIndividualCap() * stackno
	return offset + s.sizes[stackno] - 1

}

func (s stack) getIndividualCap() int {
	return s.capacity / s.numberofstacks

}
