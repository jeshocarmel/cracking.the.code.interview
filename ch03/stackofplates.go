package ch03

type plates struct {
	items     [][]int
	threshold int
}

func newSetOfPlates(threshold int) *plates {

	tmp := new(plates)
	tmp.threshold = threshold

	tmp.items = make([][]int, 0)
	tmp.items = append(tmp.items, []int{})

	return tmp
}

func (p *plates) push(i int) {

	currentStackIndex := len(p.items) - 1
	sizeOfCurrentStack := len(p.items[currentStackIndex])

	if sizeOfCurrentStack >= p.threshold {
		p.items = append(p.items, []int{})
		currentStackIndex++
	}

	p.items[currentStackIndex] = append(p.items[currentStackIndex], i)
}

func (p *plates) pop() int {

	currentStackIndex := len(p.items) - 1
	lastElementIndexInCurrentStack := len(p.items[currentStackIndex]) - 1

	deleted := p.items[currentStackIndex][lastElementIndexInCurrentStack]
	p.items[currentStackIndex] = p.items[currentStackIndex][0:lastElementIndexInCurrentStack]

	if len(p.items[currentStackIndex]) == 0 {
		p.items = p.items[0 : len(p.items)-1]
	}

	return deleted
}

func (p *plates) popArray(i int) int {

	if i < len(p.items) {
		deleted := p.items[i][len(p.items)-1]
		p.items[i] = p.items[i][0 : len(p.items)-1]
		return deleted
	}

	return 0
}
