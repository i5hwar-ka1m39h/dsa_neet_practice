package sol013

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{2, 3, 10, 5, 7, 8, 9}, 36},
	}

	for i, tc := range tests {
		got := MaxArea(tc.height)
		if got != tc.expected {
			t.Errorf("Test %d: MaxArea(%v) = %d; want %d", i+1, tc.height, got, tc.expected)
		}
	}
}
