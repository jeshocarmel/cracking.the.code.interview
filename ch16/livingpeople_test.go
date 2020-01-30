package ch16

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestLivingPeople(t *testing.T) {

	raj := Person{birth: 1948, death: 1989}
	suresh := Person{birth: 1978, death: 1967}
	ramesh := Person{birth: 1989, death: 2000}
	diksha := Person{birth: 1969, death: 1999}
	raju := Person{birth: 1956, death: 1989}
	ranjith := Person{birth: 1999, death: 1999}
	suruli := Person{birth: 1999, death: 2000}

	list := []Person{raj, suresh, ramesh, diksha, raju, ranjith, suruli}

	ans := maxLivingPeopleYearBF(list)
	assert.Equal(t, ans, 1989)

	ans = maxLivingPeopleOptimal(list)
	assert.Equal(t, ans, 1989)
}
