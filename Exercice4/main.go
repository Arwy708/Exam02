package main

import "fmt"

func CountLetters(f func(string) bool, tab []string) int {
	count := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(CountLetters(func(s string) bool {
		return s != " "
	}, []string{"Hello", " ", "World", "!", "."}))
}
