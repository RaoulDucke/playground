package main

import "fmt"

type empty struct{}

func deleteDupl(a []int) []int {
	res := make([]int, 0)
	m := make(map[int]struct{})
	for _, v1 := range a {
		_, ok := m[v1]
		if !ok {
			m[v1] = empty{}
			res = append(res, v1)
		}
	}
	return res
}

func main() {
	res := deleteDupl([]int{2, 2, 2, 1, 2, 2, 2, 7, 7, 6, 6, 12, 12, 2})
	fmt.Printf("%v \n", res)
}
