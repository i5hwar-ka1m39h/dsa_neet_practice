package sol007

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		input []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2 with zero",
			input:    []int{0, 1, 2, 3},
			expected: []int{6, 0, 0, 0},
		},
		{
			name:     "All ones",
			input:    []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{1},
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result := ProductExceptSelf(test.input)

			if !reflect.DeepEqual(result, test.expected){
				t.Error("Expected outcome", test.expected, " the output is :", result)
			}
		})
	}

}