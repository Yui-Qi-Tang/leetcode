package sol

import (
	"testing"
)

func TestSol(t *testing.T) {
	testcases := []struct {
		s       string
		induces []int
		want    string
	}{
		{s: "codeleet", induces: []int{4, 5, 6, 7, 0, 2, 1, 3}, want: "leetcode"},
		{s: "abc", induces: []int{0, 1, 2}, want: "abc"},
		{s: "art", induces: []int{1, 0, 2}, want: "rat"},
	}

	for _, ts := range testcases {
		result := restoreString(ts.s, ts.induces)
		if result != ts.want {
			t.Fatal("expected:", ts.want, "but got:", result)
		}
	}

	t.Log("passed")
}
