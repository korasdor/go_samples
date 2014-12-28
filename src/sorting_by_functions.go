package main

import (
	"fmt"
	"sort"
)

type ByLength []string

/**
* Длинна массива
**/
func (s ByLength) Len() int {
	return len(s)
}

/**
* Меняем местами элементы массива
**/
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

/**
* Сравниваем элементы массива
**/
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
