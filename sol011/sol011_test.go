package sol011

import (
	"reflect"
	"testing"
)


func TestTwoSum(t *testing.T){
	tests := []struct{
		name string
		input_arr []int
		target int
		expected []int
	}{
		{
			name:"contains the target ",
			input_arr: []int{2,7,9,10},
			target: 9,
			expected: []int{1,2},
		},

		{
			name:"contains the target ",
			input_arr: []int{2,7,9,10},
			target: 9,
			expected: []int{1,2},
		},
	
		{
			name:"contains the target ",
			input_arr: []int{2,7,9,10},
			target: 9,
			expected: []int{1,2},
		},	
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			output := TwoSum(test.input_arr, test.target)
			if !reflect.DeepEqual(output, test.expected){
				t.Error("we expected:", test.expected, "but we got:", output)
			}
		})
	}
}