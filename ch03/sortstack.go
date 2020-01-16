package ch03

//stack implementation starts here
type sortstack struct {
	items []int
}

func newSortStack() *sortstack {

	s := sortstack{}
	s.items = make([]int, 0)
	return &s
}

func (s *sortstack) push(i int) {

	s.items = append(s.items, i)
}

func (s *sortstack) pop() int {

	deleted := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return deleted

}

func (s *sortstack) peek() int {
	return s.items[len(s.items)-1]
}

func (s *sortstack) isEmpty() bool {
	return len(s.items) == 0
}

//stack implementation end here

//sort method
func (s *sortstack) sort() {

	r := newSortStack()

	for !s.isEmpty() {

		tmp := s.pop()

		for !r.isEmpty() && r.peek() > tmp {
			s.push(r.pop())
		}

		r.push(tmp)
	}

	for !r.isEmpty() {
		s.push(r.pop())
	}
}
