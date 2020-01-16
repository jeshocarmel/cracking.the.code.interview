package ch03

import "testing"
import "github.com/magiconair/properties/assert"

func TestQueueViaStack(t *testing.T) {
	myqueue := newCustomQueue()
	myqueue = myqueue.add(1)
	myqueue = myqueue.add(2)
	myqueue = myqueue.remove()
	myqueue = myqueue.add(3)
	myqueue = myqueue.add(4)
	myqueue = myqueue.remove()
	myqueue = myqueue.add(5)

	assert.Equal(t, myqueue.list(), []int{3, 4, 5})
	assert.Equal(t, myqueue.peek(), 5)

	myqueue = myqueue.remove()

	assert.Equal(t, myqueue.list(), []int{4, 5})
	assert.Equal(t, myqueue.peek(), 5)

}
