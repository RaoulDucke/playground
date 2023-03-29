package main

func indexOf(a []int, b int) int {

	for i, v := range a {
		if v == b {
			return i
		}

	}
	return -1
}

func main() {

	// Написать функцию, которая принимает слайс интов, 2 параметром принимает инт значение, которое нужно найти в слайсе. На выходе возвращаем индекс этого значения.
	// indexOf([]int{1, 2, 3}, 3). Должно вернуться 2.
	res1 := indexOf([]int{1, 2, 3}, 3)
	print(res1)
	print("\n")
	res2 := indexOf([]int{}, 3)
	print(res2)
	print("\n")
	res3 := indexOf([]int{2, 3, 1, 4, 7, 12, 15}, 2)
	print(res3)
	print("\n")
	res4 := indexOf([]int{1, 2, 2}, 2)
	print(res4)
	print("\n")
}
