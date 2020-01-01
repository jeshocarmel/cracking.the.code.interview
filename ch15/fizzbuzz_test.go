package ch15

import "testing"

import "reflect"

func TestFizzBuzz(t *testing.T) {
	singleThreadResult := FizzBuzzSingleThreaded(20)
	//t.Log(singleThreadResult)
	routinesResult := FizzBuzzRoutines(20)
	//t.Log(routinesResult)

	eq := reflect.DeepEqual(singleThreadResult, routinesResult)
	if !eq {
		t.Fail()
	}
}
