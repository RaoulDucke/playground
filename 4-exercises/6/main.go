package main

import "fmt"

func sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				res := a[i]
				a[i] = a[j]
				a[j] = res
			}
		}
	}
	return a
}

func main() {
	res := sort([]int{2, 1, 1, 3, 1, 16, 3131, 11})

	fmt.Printf("%v \n", res)
}
