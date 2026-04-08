package sol009

func LongestConsecutive(nums []int) int {
	store := make(map[int]struct{})

	for _, i := range nums {
		store[i] = struct{}{}
	}

	longest := 0
	for num := range store {
		if _, exist := store[num-1]; !exist {
			length := 1
			for {
				if _, exist := store[num+length]; exist {
					length++
				} else {
					break
				}
			}

			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
