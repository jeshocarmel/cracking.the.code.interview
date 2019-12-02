package ch10

import "testing"

func TestGroupAnagrams(t *testing.T) {

	arr := []string{"dust", "acre", "race", "fogo", "care"}
	groupedarr := groupAnagrams(arr)
	resultarr := []string{"dust", "acre", "race", "care", "fogo"}

	if !isEqual(groupedarr, resultarr) {
		t.Fail()
	}

}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
