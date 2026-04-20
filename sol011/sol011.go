package sol011






func TwoSum(numbers []int, target int) []int {
	arr_length := len(numbers)

	left := 0
	right := arr_length - 1


	for left < right{
		sum := numbers[left]+ numbers[right]
		if sum > target{
			right--
		}else if sum < target{
			left++
		}else{
			return []int{left+1, right+1}
		}
	}


	return []int{}
	
	
    
}