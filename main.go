package main

import (
	. "fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	c := make(chan int)
	// c := make(chan int, 1) // 비동기 채널 버퍼링. main 함수는 첫번째 고루틴이다.
	// c <- 1
	go func() {
		c <- 1
	}()
	msg := <-c

	Println(msg)

	// Sort 구현
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
		{"John", 8},
	}

	sort.Sort(ByName(kids))
	Println(kids)
	sort.Sort(ByAge(kids))
	Println(kids)
}
