package piscine

func Capitalize(s string) string {
	letters := []byte(s)
	for i := 0; i < len(letters); i++ {
		letters[i] = ToLower(letters[i])
	}

	for i := 0; i < len(letters)-1; i++ {
		if !IsAlphaNumeric(letters[i]) {
			letters[i+1] = ToUpper(letters[i+1])
		}
	}

	letters[0] = ToUpper(letters[0])

	return string(letters)

}

func IsAlphaNumeric(r byte) bool {
	result := (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z')
	return result
}

func ToLower(b byte) byte {
	if isUpper(b) {
		return b + 32
	}
	return b
}

func ToUpper(b byte) byte {
	if isLower(b) {
		return b - 32
	}
	return b
}

func IsUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isLower(b byte) bool {
	return b >= 'a' && b <= 'z'

}
