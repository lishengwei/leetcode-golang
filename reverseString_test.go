package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := []byte("hello")
	output := reverseString(input)
	want := []byte("olleh")
	if !reflect.DeepEqual(output, want) {
		t.Errorf("error")
	}
}
