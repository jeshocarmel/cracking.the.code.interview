package ch10

import (
	"math/rand"
	"testing"
	"time"
)

func TestGenerateMatrix(t *testing.T) {
	mat := GenerateMatrix(4, 4)

	for i := 0; i+1 < len(mat); i++ {
		for j := 0; j+1 < len(mat[0]); j++ {
			if mat[i+1][j] < mat[i][j] {
				t.Error("condition fail in column asc sorting")
				t.Fail()
			}

			if mat[i][j+1] < mat[i][j] {
				t.Error("condition fail in row asc sorting")
				t.Fail()
			}
		}
	}
}

func TestSearchMatrix(t *testing.T) {

	mat := GenerateMatrix(4, 4)
	rand.Seed(time.Now().UnixNano())
	sRow := rand.Intn(4)
	sCol := rand.Intn(4)

	isFound := SearchMatrix(mat, mat[sRow][sCol])

	if !isFound {
		t.Errorf("unable to find")
		t.Fail()
	}

	isFound = SearchMatrix(mat, 100)

	if isFound {
		t.Errorf("invalid entry found")
		t.Fail()
	}

}
