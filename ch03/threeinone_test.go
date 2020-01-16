package ch03

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestThreeInOne(t *testing.T) {

	// creating three stacks numbered as 0,1,2 with a combined capacity of 12. each stack gets 4 as fixed allocation
	mystack := newStack(3, 12)
	mystack.push(0, 1)
	mystack.push(1, 2)
	mystack.push(2, 3)

	assert.Equal(t, mystack.items, []int{1, 0, 0, 0, 2, 0, 0, 0, 3, 0, 0, 0})

	mystack.push(0, 4)
	mystack.push(1, 5)

	assert.Equal(t, mystack.items, []int{1, 4, 0, 0, 2, 5, 0, 0, 3, 0, 0, 0})

	mystack.pop(0)
	assert.Equal(t, mystack.items, []int{1, 0, 0, 0, 2, 5, 0, 0, 3, 0, 0, 0})

	mystack.pop(2)
	assert.Equal(t, mystack.items, []int{1, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0, 0})

	mystack.push(2, 6)
	assert.Equal(t, mystack.items, []int{1, 0, 0, 0, 2, 5, 0, 0, 6, 0, 0, 0})

}
