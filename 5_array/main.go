package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	//var arr [2]int
	//arr := [...]int{1, 2, 3}

	for i, val := range arr {
		fmt.Println(i, val)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	nums := []int{1, 2, 3, 4, 5}

	names := []string{"Alice", "Bob", "Charlie"}

	s := arr[1:3] // [2 3]
	s2 := arr[:3] // [1 2 3]
	s3 := arr[3:] // [4 5]
	s4 := arr[:]  // весь

	s := []int{1, 2, 3}
	s = append(s, 4, 5)    // [1 2 3 4 5]
	s = append(s, 6, 7, 8) // [1 2 3 4 5 6 7 8]

	t := []int{9, 10}
	s = append(s, t...)

	src := []int{1, 2, 3}
	dst := make([]int, 2)
	n := copy(dst, src) //  [1 2]

	// срезы это ссылки на массив
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[0:3] // [1 2 3]
	s2 := arr[2:5] // [3 4 5]
	s1[2] = 99
	fmt.Println(s2) // [99 4 5]

	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Println(i, v)
	}
	//без индекса
	for _, v := range nums {
		fmt.Println(v)
	}

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix = append(matrix, []int{10, 11, 12})

}
