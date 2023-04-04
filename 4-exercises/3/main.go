package main

import "fmt"

func deleteValue(a []int, b int) []int {
	res := make([]int, 0)
	for _, v := range a {
		if b == v {
			continue
		}
		res = append(res, v)
	}
	return res
}

func main() {
	res := deleteValue([]int{1, 2, 3}, 2)
	fmt.Printf("%v\n", res)
}
