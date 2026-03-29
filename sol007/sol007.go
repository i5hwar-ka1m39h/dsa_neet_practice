package sol007



func ProductExceptSelf(nums []int) []int {
    length_of_arr := len(nums)

	pre:=make([]int, length_of_arr)
	post:= make([]int, length_of_arr)

	var result []int
	pre[0] = 1
	post[length_of_arr - 1] = 1

	for i := 1; i < length_of_arr ; i++{
		pre[i] = nums[i - 1] * pre[i -1 ]

	}

	for i := length_of_arr -2; i >= 0; i--{
		post[i] = nums[i+1] * post[i+1]
	}

	for i := 0; i < length_of_arr; i++{
		result = append(result, pre[i]*post[i])
	}
	return  result
}