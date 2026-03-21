package sol003


func TwoSum(nums []int, target int) []int {
	store := make(map[int]int)
	var output []int

	for i, val := range nums {
		if index, ok := store[target-val]; ok {
			output = append(output, i, index)
		}
		store[val] = i
	}
	if len(output) == 0 {
		return []int{}
	}
	return output
}
