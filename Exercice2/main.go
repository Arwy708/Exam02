package main

import "fmt"

func concat(arg ...string) string {
	result := ""
	for _, args := range arg {
		result += args + "\n"
	}
	return result
}
func main() {
	fmt.Println(concat("Hello", "World", "!"))
}
