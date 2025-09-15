package main

import "fmt"

func isSorted(a []int, f func(int, int) bool) bool {
	ascending := true
	descending := true
	for i := 0; i < len(a)-1; i++ {
		if !f(a[i], a[i+1]) {
			ascending = false
		}
		if f(a[i], a[i+1]) {
			descending = false
		}
	}
	return ascending || descending
}
func f(a1, a2 int) bool {
	return a1 <= a2
}
func main() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{5, 4, 3, 2, 1}
	a3 := []int{1, 3, 2, 4, 5}

	result1 := isSorted(a1, f)
	result2 := isSorted(a2, f)
	result3 := isSorted(a3, f)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
