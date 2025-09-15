package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	var result []string
	char := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			char += string(s[i])

		} else if char != "" {
			result = append(result, char)
			char = ""
		}
	}
	if char != "" {
		result = append(result, char)
	}
	return result
}
func main() {
	fmt.Println(SplitWhiteSpaces("Hello World !"))
}
