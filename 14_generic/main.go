package main

import "fmt"

func First[T any](s []T) T {
	return s[0]
}

func main() {
	ints := []int{1, 2, 3}
	fmt.Println(First(ints)) // 1

	strs := []string{"a", "b"}
	fmt.Println(First(strs)) // "a"

	nums := []int32{1, 2, 3}
	nums2 := []uint32{4, 5, 6}
	fmt.Println(summa(nums))
	fmt.Println(summa(nums2))
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

type Integer32 interface {
	int32 | uint32
}

func summa[T Integer32](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
