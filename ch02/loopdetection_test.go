package ch02

import (
	"github.com/magiconair/properties/assert"
	"math/rand"
	"testing"
)

func TestDetectLoop(t *testing.T) {

	head := createLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	collisionpoint := head
	tail := head

	for i := 0; i < rand.Intn(8); i++ {
		collisionpoint = collisionpoint.next
	}

	// go the tail of the created list
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = collisionpoint

	/* uncomment this statement to view the circular loop

	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
	*/

	ans := detectLoop(head)
	assert.Equal(t, ans, collisionpoint)

	head = createLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	ans = detectLoop(head)
	if ans != nil {
		t.Fail()
	}
}
