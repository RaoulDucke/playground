package main

import "fmt"

func deleteDupl(a []int) []int {
	res := make([]int, 0)

	for i, v1 := range a {
		isDuplicate := false
		for j := i + 1; j < len(a); j++ {
			if v1 == a[j] {
				isDuplicate = true
			}
		}
		if !isDuplicate {
			res = append(res, v1)
		}
	}
	return res
}

func main() {
	res := deleteDupl([]int{2, 2, 2, 1, 2, 2, 2, 7, 7, 6, 6, 12, 12, 2})
	fmt.Printf("%v \n", res)
}
