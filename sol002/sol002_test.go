package sol002

import "testing"



func TestIsAnagram(t *testing.T){
	tests:=[]struct{
		name string
		s string
		t string
		expected bool
	}{
		{"this is anagram", "anagram", "nagaram", true},
		{"this is not an anagram", "rat", "cat", false},
		{"both are empty string", "", "", true},
		{"one is empty", "", "mean", false},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result := IsAnagram(test.s, test.t)

			if result != test.expected{
				t.Errorf("Expected %v got %v", test.expected, result)
			}
		})
	}
}