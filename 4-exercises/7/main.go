package main

import "fmt"

func maxValue(a []int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

func main() {
	res := maxValue([]int{5, 1, 2, 3, 4, 4, 4, 8, 2, 3})
	fmt.Printf("%d \n", res)
}
