package sol005

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example1",
			nums:     []int{4, 1, -1, 2, -1, 2, 3},
			k:        2,
			expected: []int{-1, 2}, // order may vary if frequencies are equal
		},
		{
			name:     "AllUnique",
			nums:     []int{1, 2, 3, 4},
			k:        2,
			expected: []int{3, 4}, // any two elements, since all have same frequency
		},
		{
			name:     "SingleElement",
			nums:     []int{5, 5, 5},
			k:        1,
			expected: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TopKFrequent(tt.nums, tt.k)
			if tt.name == "AllUnique" {
				// For AllUnique, any k unique elements from nums are valid
				if len(result) != tt.k {
					t.Errorf("TopKFrequent(%v, %d) = %v; want length %d", tt.nums, tt.k, result, tt.k)
				}
				unique := make(map[int]bool)
				for _, v := range result {
					unique[v] = true
				}
				if len(unique) != tt.k {
					t.Errorf("TopKFrequent(%v, %d) = %v; elements not unique", tt.nums, tt.k, result)
				}
				for _, v := range result {
					found := false
					for _, n := range tt.nums {
						if v == n {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("TopKFrequent(%v, %d) = %v; element %d not in input", tt.nums, tt.k, result, v)
					}
				}
			} else {
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("TopKFrequent(%v, %d) = %v; want %v", tt.nums, tt.k, result, tt.expected)
				}
			}
		})
	}
}
