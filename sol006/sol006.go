package sol006

import (
	
	"strconv"
)



func Encode(strs []string) string {
	var encode_string string
	for _, val := range strs{
		encode_string += strconv.Itoa(len(val))  + "#"  + val
	}

	return encode_string
		
	
}

func Decode(encoded string) []string {
	if len(encoded) == 0{
		return []string{}
	}
	var store []string
	i:=0
	for i < len(encoded){
		j:=i

		for string(encoded[j]) != "#"{
			j++
		}

		len_of_string, _ := strconv.Atoi(encoded[i:j])
		// if err != nil {
		// 	log.Fatalln("error converting string to number")
		// 	break
		// }

		start_of_word := j+1
		end_of_word := start_of_word+len_of_string
		store = append(store, string(encoded[start_of_word:end_of_word]))

		i = end_of_word

	}
	return store
	
}
