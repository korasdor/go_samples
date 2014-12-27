package main

import "fmt"

type rect struct {
	width, height int
}

/**
/* со ссылкой на стракт
**/
func (r *rect) area() int {
	return r.width * r.height
}

/**
/* копирует экземляр стракт
**/
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("perim:", r.perim())
	fmt.Println("area:", r.area())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim", rp.perim())
}
