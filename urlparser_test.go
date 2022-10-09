package urlparser

import (
	"testing"
)

func TestGetGoWordCountByUrls(t *testing.T) {
	testcases := []struct {
		in   []string
		want int
	}{
		{[]string{"https://golang.org", "https://golang.org", "https://golang.org", "https://golang.org",
			"https://golang.org", "https://golang.org", "https://golang.org", "https://golang.org"}, 2},
	}
	for _, tc := range testcases {
		result := GetGoWordCountByUrls(tc.in)
		if result != tc.want {
			t.Errorf("Total count: %q for %q, want %q", result, tc.in, tc.want)
		}
	}
}
