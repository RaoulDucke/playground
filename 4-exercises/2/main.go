package main

import "fmt"

func indexOf(a []int, b int) []int {
	res := make([]int, 0, 0)

	for i, v := range a {
		if b == v {
			res = append(res, i)
		}
	}
	return res
}

func main() {

	res := indexOf([]int{1, 3, 4, 5, 4, 16, 424, 4, 4, 4}, 4)
	fmt.Printf("%v \n", res)

}
