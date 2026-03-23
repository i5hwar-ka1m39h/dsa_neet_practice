package sol003

import (
	"reflect"
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{"exists", []int{2, 7, 3, 4}, 9, []int{0,1}},
		{"double exist", []int{1, 3,  2, 45, 9}, 4, []int{ 0,1}},
		{"no solution", []int{1, 2, 3}, 10, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TwoSum(test.input, test.target)
			slices.Sort(result)
			
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v got %v", test.expected, result)
			}
		})
	}
}
