package sol001

import "testing"


func TestContainsDuplicate(t *testing.T){
	tests:= []struct{
		name string
		input []int
		expected bool
	}{
		{"duplicate exist", []int{1, 2,3,4,2}, true},
		{"doesn't contains the duplicate", []int{1,3,4,0,2,45,9}, false},
		{"empty arrays", []int{}, false},
		{"contains double duplcate", []int{1,2,3,2,4,5,4}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result:= ContainsDuplicate(test.input)

			if result != test.expected{
				t.Errorf("expected %v got %v", test.expected, result)
			}
		})
	}
}