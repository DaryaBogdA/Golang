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
		fmt.Println("not an int", n) // not an int
	}

	player := Player{34, "dsfdg"}
	fmt.Println(player)
	player.Reset()
	Reset(&player) // 2 variant
	fmt.Println(player)

	describe("4")
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

func describe(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Целое: %d\n", v)
	case string:
		fmt.Printf("Строка: %q\n", v)
	case bool:
		fmt.Printf("Логическое: %t\n", v)
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}

type Whispper interface {
	Whisp() string
}

type Yeller interface {
	Yell() string
}

type Talker interface {
	Whispper
	Yeller
}

func talk(t Talker) {
	fmt.Println(t.Whisp())
	fmt.Println(t.Yell())
}
