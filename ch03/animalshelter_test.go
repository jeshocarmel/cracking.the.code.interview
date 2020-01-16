package ch03

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAnimalShelter(t *testing.T) {

	as := new(AnimalShelter)
	as.enqueue("Brucy", dog)
	as.enqueue("Ginger", cat)
	as.enqueue("Pepper", cat)
	as.enqueue("Berry", cat)
	as.enqueue("Tommy", dog)
	as.enqueue("Phil", dog)

	assert.Equal(t, as.dequeueAny().name, "Brucy")
	assert.Equal(t, as.dequeueCat().name, "Ginger")
	assert.Equal(t, as.dequeueDog().name, "Tommy")
	assert.Equal(t, as.dequeueAny().name, "Pepper")

}
