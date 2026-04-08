package sol009

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct{
		name string
		input []int
		expected int
	}{
		{
			name:"empty array",
			input: []int{},
			expected: 0,

		},
		{
			name:"single length array",
			input: []int{5},
			expected: 1,
		},
		{ 
			name: "only single sequence",
			input: []int{100, 200, 3000},
			expected: 1,
		},
		{
			name:"some sequence with lenght",
			input: []int{2,3,4,1,100, 200, 101,102, 201},
			expected: 4,
		},{
			name: "sorted sequence",
			input: []int{1,2,3,4,5,6,7,8},
			expected: 8,
		},
	}

	for _, test := range tests{
		t.Run(test.name,func(t *testing.T) {
			result :=LongestConsecutive(test.input)

			if result != test.expected{
				t.Error("expected output is :", test.expected, " we got :", result)
			}
		})
	}
}