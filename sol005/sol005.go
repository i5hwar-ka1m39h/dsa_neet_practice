package sol005

import (
	"fmt"
	"sort"
)

func TopKFrequent(nums []int, k int) []int{
	store := make(map[int]int)

	var result []int
	var tempStorage [][]int

	for _, val := range nums {
		if _, ok := store[val]; ok {
			store[val]++
		} else {

			store[val] = 1
		}
	}

	for k, v := range store {
		tempStorage = append(tempStorage, []int{k, v})
	}

   sort.Slice(tempStorage, func(i, j int) bool {
      return tempStorage[i][1] > tempStorage[j][1]
   })
	fmt.Println(tempStorage)

   for i:=0; i < k; i++{
      result = append(result, tempStorage[i][0])
   }
   return result

}
