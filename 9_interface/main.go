package main

import "fmt"

type MyInterface interface {
	func1()
	func2(x int) int
}
type MyInt int

func (m MyInt) func1() {}
func (m MyInt) func2(x int) int {
	return x * x
}

var empt any = "hello"

func main() {
	var i MyInt
	fmt.Println(i.func2(5))

	m := MyInt(1)
	execute(m)

	// пустой интерфейс

	у, ok := empt.(string)
	if ok {
		fmt.Println(у) // hello
	}

	n, ok := empt.(int)
	if !ok {
		fmt.Println("not an int") // not an int
	}

	player := Player{34, "dsfdg"}
	fmt.Println(player)
	player.Reset()
	Reset(&player) // 2 variant
	fmt.Println(player)
}

type Resetter interface {
	Reset()
}

func Reset(r Resetter) { // 2 variant
	r.Reset()
}

type Player struct {
	health   int
	position string
}

func (p *Player) Reset() {
	p.health = 0
	p.position = "dsf"
}

func execute(i MyInt) {
	i.func1()
}
