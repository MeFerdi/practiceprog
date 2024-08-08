package main

func RepeatAlpha(s string) string {
	result := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			index := int(char - 'a' + 1)
			for i := 0; i < index; i++ {
				result += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			index := int(char - 'A' + 1)
			for i := 0; i < index; i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}

	return result
}
