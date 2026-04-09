package sol010

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re:= regexp.MustCompile("[^a-zA-Z0-9]+")

	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	length_str := len(s)

	left := 0
	right := length_str -1
	
	for left <= right{
		if s[left] != s[right]{
			return false;
		}
		left++
		right--
	}

	return true
    
}