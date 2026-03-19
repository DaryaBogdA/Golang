package main

import "fmt"

const (
	s0 = iota // + 9
	s1
	s2
	s3
	_
	s5
)

type Direction byte

func main() {
	north := South
	fmt.Println(north)
}

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string { // елси переменная south результат south
	switch d {
	case North:
		return fmt.Sprintf("North")
	case South:
		return fmt.Sprintf("South")
	case East:
		return fmt.Sprintf("East")
	case West:
		return fmt.Sprintf("West")
	default:
		return "othrer"
	}
}

func (d Direction) String() string { // елси переменная south результат east
	return []string{"North", "South", "East", "West"}[d]
}
