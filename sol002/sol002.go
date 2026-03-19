package sol002

import (
	"slices"
	"strings"
)

func IsAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	sArr:= strings.Split(s, "")
	slices.Sort(sArr)
	sortedS:=strings.Join(sArr, "")

	tArr := strings.Split(t, "")
	slices.Sort(tArr)
	sortedT := strings.Join(tArr,"")


	if sortedS == sortedT{
		return true
	}
   return false 
}