package sol012

import "sort"




func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res:= [][]int{}


	for i := 0; i < len(nums); i++{
		a := nums[i]

		if a > 0{
			break
		}

		if i > 0 && a == nums[i - 1]{
			continue
		}

		l:= i+1
		r := len(nums) -1

		for l < r{
			sum_three := a + nums[l] + nums[r]

			if sum_three > 0 {
				r--
			}else if sum_three < 0 {
				l++
			}else{
				res = append(res,[]int{a, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1]{
					l++
				}
			}

		}
	}

	return  res
    
}