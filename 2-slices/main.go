package main

import "fmt"

func main() {
	// 	// Массивы
	// 	a1 := [3]int{
	// 		1, 2, 3,
	// 	}
	// 	for i := 0; i < len(a1); i++ {
	// 		println(fmt.Sprintf("%d -%d", i, a1[i]))
	// 	}
	// 	for i, v := range a1 {
	// 		println(fmt.Sprintf("%d -%d", i, v))

	// 	}
	// 	a2 := [5]int{}
	// 	a2[0] = 4
	// 	a2[1] = 5
	// 	a2[3] = 6
	// 	for i, v := range a2 {
	// 		println(fmt.Sprintf("%d -%d", i, v))
	// 	}
	// Слайсы
	s1 := []int{
		1, 3, 2,
	}
	println(fmt.Sprintf("%d -%d", cap(s1), len(s1)))
	s1 = append(s1, 4)
	println(fmt.Sprintf("%d -%d", cap(s1), len(s1)))
	s2 := make([]int, 0, 12)
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	println(fmt.Sprintf("%d -%d", cap(s2), len(s2)))
	s3 := make([]int, 3, 3)
	s3 = append(s3, 1)
	s3 = append(s3, 2)
	for i, v := range s3 {
		println(fmt.Sprintf("%d -%d", i, v))
	}
}
