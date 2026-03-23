package sol004

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:  "example1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name:  "single",
			input: []string{"abc"},
			expected: [][]string{
				{"abc"},
			},
		},
		{
			name:     "empty",
			input:    []string{},
			expected: [][]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GroupAnagrams(test.input)
			// Sort inner slices and outer slice for comparison
			for _, group := range result {
				slices.Sort(group)
			}
			for _, group := range test.expected {
				slices.Sort(group)
			}
			slices.SortFunc(result, func(a, b []string) int {
				if len(a) == 0 || len(b) == 0 {
					return len(a) - len(b)
				}
				return strings.Compare(a[0], b[0])
			})
			slices.SortFunc(test.expected, func(a, b []string) int {
				if len(a) == 0 || len(b) == 0 {
					return len(a) - len(b)
				}
				return strings.Compare(a[0], b[0])
			})
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
