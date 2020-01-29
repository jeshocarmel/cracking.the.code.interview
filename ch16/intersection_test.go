package ch16

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIntersection(t *testing.T) {

	start1 := point{x: 2.35, y: 0}
	end1 := point{x: 4, y: 10}

	start2 := point{x: 1, y: 0}
	end2 := point{x: 8, y: 14}

	intersectionPoint := intersection(start1, end1, start2, end2)

	assert.Equal(t, intersectionPoint.x, float32(3.0149252))
	assert.Equal(t, intersectionPoint.y, float32(4.0298505))

}
