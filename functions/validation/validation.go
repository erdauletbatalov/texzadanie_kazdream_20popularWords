package validation

import (
	"kazdream/functions/bytes"
)

func CheckValid(text []byte) bool {
	if bytes.Equal(text, []byte{}) == 0 {
		return true
	}
	for _, val := range text {
		if bytes.IsAscii(val) {
			return false
		}
	}
	return true
}
