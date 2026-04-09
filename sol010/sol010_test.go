package sol010

import "testing"


func TestIsPalindromeFn(t *testing.T){
	tests := []struct{
		name string
		input string
		expected bool
	}{
		{
			name:"is valid palidrome",
			input: "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:"not a valid palidrome",
			input:"my cat ate shit like a dog",
			expected: false,
		},
		{
			name:"empty string input",
			input: "",
			expected: true,
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			output := isPalindrome(test.input)
			if output != test.expected{
				t.Error("we expected:", test.expected, "but we got:", output)
			}
		})
	}
}