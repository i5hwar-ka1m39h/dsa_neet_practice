package sol004

import (
	"slices"
	"strings"
)


func GroupAnagrams(strs []string) [][]string {
	
	if len(strs) == 0{
		return [][]string{}
	}
	store:=make(map[string][]string)

	var output [][]string
   for _,word := range strs{
		chars := strings.Split(word, "")
		slices.Sort(chars)
		sortedWord := strings.Join(chars, "")
		if wordArr, ok := store[sortedWord];ok{
			wordArr = append(wordArr, word)
		}

		store[sortedWord]= append(store[sortedWord], word)
   } 

   for _, val := range store{
	output = append(output, val)
   }

   return output
}