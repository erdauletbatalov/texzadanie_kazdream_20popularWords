package bytes

func IsAscii(char byte) bool {
	if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
		return true
	}
	return false
}

func ToLower(word []byte) []byte {
	hasUpper := false
	for i := 0; i < len(word); i++ {
		char := word[i]
		hasUpper = ('A' <= char && char <= 'Z')
		if hasUpper {
			break
		}
	}
	if !hasUpper {
		return word
	}
	result := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		loweredChar := word[i]
		if 'A' <= loweredChar && loweredChar <= 'Z' {
			loweredChar += 'a' - 'A'
		}
		result[i] = loweredChar
	}
	return result
}

func Equal(a, b []byte) int {
	if len(a) < len(b) {
		for i, v := range a {
			if v > b[i] {
				return 1
			} else if v < b[i] {
				return -1
			}
		}
		return -1
	}

	if len(a) > len(b) {
		for i, v := range b {
			if v < a[i] {
				return 1
			} else if v > a[i] {
				return -1
			}
		}
		return 1
	}
	if len(a) == len(b) {
		for i, v := range a {
			if v > b[i] {
				return 1
			} else if v < b[i] {
				return -1
			}
		}
		return 0
	}

	return 0
}

func IntToBytes(x int) []byte {
	arr := []byte{}
	for x != 0 {
		arr = append(arr, byte(x%10+48))
		x /= 10
	}
	return reverse(arr)
}

func reverse(a []byte) []byte {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
