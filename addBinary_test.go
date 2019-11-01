package leetcode

import "testing"

func TestAddBinary(t *testing.T) {
	a := "1111"
	b := "1111"
	got := addBinary(a, b)
	wang := "11110"

	if got != wang {
		t.Errorf("got %s wang %s", got, wang)
	}

}
