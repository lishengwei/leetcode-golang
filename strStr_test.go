package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	haystack := ""
	needle := ""
	got := strStr(haystack, needle)
	want := -1

	if got != want {
		t.Errorf("Error got : %d, want %d", got, want)
	}
}
