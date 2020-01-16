package ch03

import "testing"

import "github.com/magiconair/properties/assert"

func TestStackMin(t *testing.T) {

	mystack := newMinStack()
	mystack = mystack.push(7)
	mystack = mystack.push(5)
	mystack = mystack.push(3)
	mystack = mystack.push(4)
	mystack = mystack.push(1)

	minimum := mystack.min()
	assert.Equal(t, minimum, 1)

	mystack = mystack.pop()
	minimum = mystack.min()
	assert.Equal(t, minimum, 3)

	mystack = mystack.push(2)
	minimum = mystack.min()
	assert.Equal(t, minimum, 2)

}
