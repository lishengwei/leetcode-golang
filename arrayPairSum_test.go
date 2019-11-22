package leetcode

import "testing"

func TestArrayPairSum(t *testing.T) {
	input := []int{1, 4, 3, 2}

	output := arrayPairSum(input)

	want := 4
	if want != output {
		t.Errorf("error")
	}
}
