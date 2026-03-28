package sol006

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	testCases := []struct {
		input []string
		desc  string
	}{
		{[]string{"hello", "world"}, "basic_two_words"},
		{[]string{"", "a", "bc"}, "empty and short strings"},
		{[]string{"#hash", "123", "!@#"}, "special characters"},
		{[]string{"oneword"}, "single word"},
		{[]string{"", "", ""}, "all empty strings"},
		{[]string{}, "empty input slice"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
		encoded := Encode(tc.input)
		decoded := Decode(encoded)
		if !reflect.DeepEqual(tc.input, decoded) {
			t.Errorf("%s: Encode/Decode failed. Input: %v, Decoded: %v", tc.desc, tc.input, decoded)
		}

		})
	}
}
